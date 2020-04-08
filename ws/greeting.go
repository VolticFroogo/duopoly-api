package ws

import (
	"errors"
	"fmt"

	"github.com/VolticFroogo/duopoly-api/game"
	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var (
	ErrPlaying = errors.New("game not joinable as it is in progress")
)

func greeting(conn *websocket.Conn, secret string) (g *game.Game, playerID int, err error) {
	// Read the initial message from the client which should be a greeting.
	var msg message.Message
	err = conn.ReadJSON(&msg)
	if err != nil {
		_ = conn.WriteJSON(message.Message{
			Type: message.ResponseError,
			Data: fmt.Sprintf("reading json request: %s", err.Error()),
		})
		return
	}

	// Decode the message as a greeting.
	var greeting requestGreeting
	err = mapstructure.Decode(msg.Data.(map[string]interface{}), &greeting)
	if err != nil {
		_ = conn.WriteJSON(message.Message{
			Type: message.ResponseError,
			Data: fmt.Sprintf("decoding message: %s", err.Error()),
		})
		return
	}

	if g2, ok := game.Games[greeting.ID]; ok {
		// If a game exists with the ID provided, get it.
		g = g2
	} else {
		// Otherwise, create one.
		g = game.New(greeting.ID, secret)
	}

	// Find the player's ID from their secret.
	playerID = g.GetPlayerID(secret)

	// If the game is in progress, and the player isn't a part of it, reject them.
	if g.Playing && playerID == game.NullPlayerID {
		_ = conn.WriteJSON(message.Message{
			Type: message.ResponsePlaying,
		})
		err = ErrPlaying
		return
	}

	if playerID == game.NullPlayerID {
		// If the player is joining a pre-game lobby, announce it.
		playerID = g.PlayerJoined(&game.Player{
			Secret:    secret,
			WS:        conn,
			Name:      greeting.Name,
			Connected: true,
		})
	} else {
		// Otherwise, update the player's previous state to reflect this new connection.
		// TODO: this could in theory allow two concurrent writable connections for a player,
		//  but only one will receive broadcasts; fix / look into this.
		g.Players[playerID].WS = conn
		g.Players[playerID].Connected = true
	}

	// Send the player the current state of the game as a greeting.
	err = conn.WriteJSON(message.Message{
		Type: message.ResponseGreeting,
		Data: *g,
	})
	if err != nil {
		_ = conn.WriteJSON(message.Message{
			Type: message.ResponseError,
			Data: fmt.Sprintf("writing greeting: %s", err.Error()),
		})
	}

	return
}

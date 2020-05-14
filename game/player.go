package game

import (
	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gorilla/websocket"
)

type Player struct {
	ID        int             `json:"-"`
	Secret    string          `json:"-"`
	WS        *websocket.Conn `json:"-"`
	Connected bool            `json:"connected"`
	Name      string          `json:"name"`
	Position  int             `json:"position"`
	Money     int             `json:"money"`
	Sentence  int             `json:"sentence"`
	Bankrupt  bool            `json:"bankrupt"`
}

func (game *Game) PlayerJoined(player *Player) int {
	// Prevent game concurrency collisions.
	game.mutex.Lock()
	defer game.mutex.Unlock()

	// Assign the player the next ID and increment it.
	player.ID = game.nextPlayer
	game.nextPlayer++

	// Add the player to the player map.
	game.Players[player.ID] = player

	// Broadcast the new player's connection to everyone but the new player.
	game.broadcast(message.Message{
		Type: message.ResponseJoined,
		Data: *player,
	}, map[int]bool{player.ID: true})

	// Return the newly generated ID.
	return player.ID
}

func (game *Game) PlayerLeft(playerID int) {
	// Prevent game concurrency collisions.
	game.mutex.Lock()
	defer game.mutex.Unlock()

	if game.Playing {
		// If game is in progress, set the player as disconnected.
		game.Players[playerID].WS = nil
		game.Players[playerID].Connected = false
	} else {
		// Otherwise, remove the player from the game.
		delete(game.Players, playerID)
	}

	// Broadcast the player's disconnection.
	game.broadcast(message.Message{
		Type: message.ResponseLeft,
		Data: playerID,
	}, nil)
}

func (game *Game) GetPlayerID(secret string) int {
	// Find the player's ID given a secret.
	for _, player := range game.Players {
		if player.Secret == secret {
			return player.ID
		}
	}

	// If not found, return a null ID.
	return NullPlayer
}

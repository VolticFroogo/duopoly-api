package game

import (
	"fmt"

	"github.com/VolticFroogo/duopoly-api/message"
)

const (
	MaxChatLength = 2048
)

type responseChat struct {
	Player  int    `json:"player"`
	Message string `json:"message"`
}

func (game *Game) chat(msg message.Message, playerID int) {
	// Get the text from the message data.
	text := msg.Data.(string)

	// Verify the text isn't over the maximum length.
	if len(text) > MaxChatLength {
		_ = game.Players[playerID].WS.WriteJSON(message.Message{
			Type: message.ResponseError,
			Data: fmt.Sprintf("chat cannot be longer than %d", MaxChatLength),
		})
		return
	}

	// Broadcast the message to all but the sender.
	game.broadcast(message.Message{
		Type: message.ResponseChat,
		Data: responseChat{
			Player:  0,
			Message: text,
		},
	}, map[int]bool{playerID: true})
}

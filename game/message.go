package game

import (
	"github.com/VolticFroogo/duopoly-api/message"
)

func (game *Game) Message(msg message.Message, playerID int) {
	// Prevent game concurrency collisions.
	game.mutex.Lock()
	defer game.mutex.Unlock()

	// Handle the message with the relevant function based on the type.
	switch msg.Type {
	case message.RequestChat:
		game.chat(msg, playerID)
		break
	}
}

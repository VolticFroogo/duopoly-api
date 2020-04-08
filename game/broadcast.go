package game

import (
	"github.com/VolticFroogo/duopoly-api/message"
)

func (game *Game) broadcast(msg message.Message, exclude map[int]bool) {
	// Broadcast the message to everyone in the game that isn't in the exclude map.
	for _, player := range game.Players {
		if player.Connected && !exclude[player.ID] {
			_ = player.WS.WriteJSON(msg)
		}
	}
}

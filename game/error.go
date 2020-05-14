package game

import "github.com/VolticFroogo/duopoly-api/message"

func (game *Game) error(playerID int, error string) {
	_ = game.Players[playerID].WS.WriteJSON(message.Message{
		Type: message.ResponseError,
		Data: error,
	})
}

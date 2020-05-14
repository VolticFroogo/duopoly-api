package game

import (
	"github.com/VolticFroogo/duopoly-api/helper"
	"github.com/VolticFroogo/duopoly-api/message"
)

func (game *Game) tryRoll(playerID int) {
	if game.Action.Type != actionTurn {
		game.error(playerID, "cannot roll: not action turn")
		return
	}

	data := game.Action.Data.(dataTurn)

	if data.PlayerID != playerID {
		game.error(playerID, "cannot roll: not your turn")
		return
	}

	if data.HasRoll {
		game.error(playerID, "cannot roll: no roll remaining")
		return
	}

	game.roll()
}

func (game *Game) roll() {
	data := game.Action.Data.(dataTurn)

	dice := helper.RandomDice()
	roll := dice[0] + dice[1]

	game.Players[data.PlayerID].Position += roll

	if game.Players[data.PlayerID].Position >= BoardLen {
		game.Players[data.PlayerID].Position -= BoardLen
	}

	// If the roll isn't a double, prevent the player from rolling again.
	if dice[0] != dice[1] {
		data.HasRoll = false
		game.Action.Data = data
	}

	game.broadcast(message.Message{
		Type: message.ResponseRoll,
		Data: dice,
	}, nil)

	game.Board[game.Players[data.PlayerID].Position].Landed(game, data.PlayerID, roll)
}

package game

import (
	"time"

	"github.com/VolticFroogo/duopoly-api/message"
)

func (game *Game) buy(playerID int) {
	if game.Action.Type != actionPurchaseDecision {
		game.error(playerID, "cannot buy: not purchase decision action")
		return
	}

	data := game.Action.Data.(dataPurchaseDecision)

	if data.PlayerID != playerID {
		game.error(playerID, "cannot buy: not your purchase decision")
		return
	}

	purchasable := game.Board[game.Players[playerID].Position].Purchasable()

	if purchasable == nil {
		game.error(playerID, "cannot buy: tile not purchasable")
		return
	}

	if purchasable.GetOwned() {
		game.error(playerID, "cannot buy: tile owned")
		return
	}

	value := purchasable.GetValue()

	if value > game.Players[playerID].Money {
		game.error(playerID, "cannot buy: not enough money")
		return
	}

	game.Players[playerID].Money -= value
	purchasable.SetOwner(playerID)

	game.Action = game.Queue[0]
	game.Queue = game.Queue[:1]

	turnData := game.Action.Data.(dataTurn)

	turnData.Timeout = time.Now().Add(turnTimeout).Unix()

	game.Action.Data = turnData

	game.broadcast(message.Message{
		Type: message.ResponseBuy,
	}, nil)
}

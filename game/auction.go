package game

import (
	"time"

	"github.com/VolticFroogo/duopoly-api/message"
)

const (
	auctionTimeoutInitial = 10
	auctionTimeout        = 5
	auctionDefaultValue   = 10
)

type dataAuction struct {
	Property int   `json:"property"`
	Bid      Bid   `json:"bid"`
	Timeout  int64 `json:"timeout"`
}

func tickAuction(game *Game) {
	data := game.Action.Data.(dataAuction)

	if data.Timeout > time.Now().Unix() {
		return
	}

	if data.Bid.PlayerID != NullPlayer {
		game.Players[data.Bid.PlayerID].Money -= data.Bid.Value
		game.Board[data.Property].Purchasable().SetOwner(data.Bid.PlayerID)
	}

	game.Action = game.Queue[0]
	game.Queue = game.Queue[1:]

	turnData := game.Action.Data.(dataTurn)

	turnData.Timeout = time.Now().Add(turnTimeout).Unix()

	game.Action.Data = turnData

	game.broadcast(message.Message{
		Type: message.ResponseAuctionWon,
		Data: game.Action,
	}, nil)
}

func (game *Game) tryAuction(playerID int) {
	if game.Action.Type != actionPurchaseDecision {
		game.error(playerID, "cannot auction: not purchase decision action")
		return
	}

	data := game.Action.Data.(dataPurchaseDecision)

	if data.PlayerID != playerID {
		game.error(playerID, "cannot auction: not your purchase decision")
		return
	}

	game.auction(playerID)
}

func (game *Game) auction(playerID int) {
	game.Action = Action{
		Type: actionAuction,
		Data: dataAuction{
			Property: game.Players[playerID].Position,
			Bid:      defaultBid,
			Timeout:  time.Now().Add(auctionTimeoutInitial).Unix(),
		},
	}

	game.broadcast(message.Message{
		Type: message.ResponseAuction,
	}, nil)
}

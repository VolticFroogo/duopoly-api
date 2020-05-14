package game

import (
	"time"
)

type dataPurchaseDecision struct {
	PlayerID int   `json:"player"`
	Timeout  int64 `json:"timeout"`
}

func tickPurchaseDecision(game *Game) {
	data := game.Action.Data.(dataPurchaseDecision)

	if data.Timeout > time.Now().Unix() {
		return
	}

	game.auction(game.Queue[0].Data.(dataTurn).PlayerID)
}

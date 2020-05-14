package game

import (
	"time"

	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gin-gonic/gin"
)

const (
	turnTimeout = time.Second * 30
)

type dataTurn struct {
	PlayerID int   `json:"player"`
	HasRoll  bool  `json:"rolled"`
	Timeout  int64 `json:"timeout"`
}

func tickTurn(game *Game) {
	data := game.Action.Data.(dataTurn)

	if data.Timeout > time.Now().Unix() {
		return
	}

	if !data.HasRoll {
		game.roll()
		return
	}

	game.endTurn(data.PlayerID, true)
}

func (game *Game) tryEndTurn(playerID int) {
	if game.Action.Type != actionTurn {
		game.error(playerID, "cannot end turn: currently not on turn action")
		return
	}

	data := game.Action.Data.(dataTurn)
	if data.PlayerID != playerID {
		game.error(playerID, "cannot end turn: not your turn")
		return
	}

	game.endTurn(data.PlayerID, false)
}

func (game *Game) endTurn(lastPlayerID int, forced bool) {
	// Get the next valid player in the game.
	var next int
	setNext := true

	for id := range game.Players {
		if game.Players[id].Bankrupt {
			continue
		}

		if id == lastPlayerID {
			setNext = true
			continue
		}

		if setNext {
			next = id
			setNext = false
		}
	}

	// Set the current action to the next player's turn.
	game.Action = Action{
		Type: actionTurn,
		Data: dataTurn{
			PlayerID: next,
			Timeout:  time.Now().Add(turnTimeout).Unix(),
		},
	}

	// Broadcast the new turn.
	game.broadcast(message.Message{
		Type: message.ResponseSetTurn,
		Data: gin.H{
			"player": next,
			"forced": forced,
		},
	}, nil)
}

package game

import (
	"time"

	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gin-gonic/gin"
)

type dataPay struct {
	Payer   int   `json:"payer"`
	Payee   int   `json:"payee"`
	Value   int   `json:"value"`
	Timeout int64 `json:"timeout"`
}

func (game *Game) addPayment(payer int, payee int, value int) {
	game.Queue = append(game.Queue, game.Action)
	game.Action = Action{
		Type: actionPay,
		Data: dataPay{
			Payer: payer,
			Payee: payee,
			Value: value,
		},
	}
}

func tickPay(game *Game) {
	data := game.Action.Data.(dataPay)

	if data.Timeout > time.Now().Unix() {
		return
	}

	if game.Players[data.Payer].Money < data.Value {
		game.bankrupt(data.Payer, data.Payee)
	}

	game.pay(data, true)
}

func (game *Game) tryPay(playerID int) {
	if game.Action.Type != actionPay {
		game.error(playerID, "cannot pay: not pay action")
		return
	}

	data := game.Action.Data.(dataPay)

	if data.Payer != playerID {
		game.error(playerID, "cannot pay: not your pay action")
		return
	}

	if game.Players[data.Payer].Money < data.Value {
		game.error(playerID, "cannot pay: not enough money")
		return
	}

	game.pay(data, false)
}

func (game *Game) pay(data dataPay, forced bool) {
	game.Players[data.Payer].Money -= data.Value

	if data.Payee != NullPlayer {
		game.Players[data.Payee].Money += data.Value
	}

	game.broadcast(message.Message{
		Type: message.ResponsePay,
		Data: gin.H{
			"forced": forced,
		},
	}, nil)
}

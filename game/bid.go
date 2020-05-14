package game

import (
	"time"

	"github.com/VolticFroogo/duopoly-api/helper"
	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gin-gonic/gin"
)

var (
	defaultBid = Bid{
		PlayerID: NullPlayer,
		Value:    auctionDefaultValue,
	}
)

type Bid struct {
	PlayerID int `json:"player"`
	Value    int `json:"value"`
}

func (game *Game) bid(msg message.Message, playerID int) {
	// TODO: some of these errors can occur just due to client latency;
	// it would probably be best to implement a safer response than error.
	if game.Action.Type != actionAuction {
		game.error(playerID, "cannot bid: not auction action")
		return
	}

	value, ok := msg.Data.(int)
	if !ok {
		game.error(playerID, "cannot bid: data is not int")
		return
	}

	if game.Players[playerID].Money < value {
		game.error(playerID, "cannot bid: not enough money")
		return
	}

	data := game.Action.Data.(dataAuction)
	if data.Bid.PlayerID == playerID {
		game.error(playerID, "cannot bid: already winning")
		return
	}

	if value <= data.Bid.Value {
		game.error(playerID, "cannot bid: value not greater than current bid")
		return
	}

	// Set the timeout to whichever is higher: the previous timeout or now plus the timeout.
	// This is necessary as the initial timeout is extended to give time to think.
	data.Timeout = helper.MaxInt64(data.Timeout, time.Now().Add(auctionTimeout).Unix())

	data.Bid.Value = value
	data.Bid.PlayerID = playerID

	game.Action.Data = data

	game.broadcast(message.Message{
		Type: message.ResponseBid,
		Data: gin.H{
			"player": playerID,
			"value":  value,
		},
	}, nil)
}

package game

import (
	"github.com/VolticFroogo/duopoly-api/message"
)

func (game *Game) Message(msg message.Message, playerID int) {
	// Prevent game concurrency collisions.
	game.mutex.Lock()
	defer game.mutex.Unlock()

	// TODO: possibly refactor this to use registered handles over a switch.
	// Handle the message with the relevant function based on the type.
	switch msg.Type {
	case message.RequestChat:
		game.chat(msg, playerID)
		break

	case message.RequestStart:
		game.start(playerID)
		break

	case message.RequestRoll:
		game.tryRoll(playerID)
		break

	case message.RequestEndTurn:
		game.tryEndTurn(playerID)
		break

	case message.RequestBuy:
		game.buy(playerID)
		break

	case message.RequestAuction:
		game.tryAuction(playerID)
		break

	case message.RequestBid:
		game.bid(msg, playerID)
		break

	case message.RequestPay:
		game.tryPay(playerID)
		break

	case message.RequestBuild:
		// TODO: implement building.
		break

	case message.RequestDemolish:
		// TODO: implement demolishing.
		break

	case message.RequestMortgage:
		// TODO: implement mortgaging.
		break

	case message.RequestUnmortgage:
		// TODO: implement unmortgaging.
		break

	case message.RequestBankrupt:
		game.tryBankrupt(playerID)
		break
	}
}

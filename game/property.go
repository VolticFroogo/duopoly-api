package game

import "time"

type Property struct {
	Value      int    `json:"-"`
	Set        int    `json:"-"`
	Rent       [6]int `json:"-"`
	HousePrice int    `json:"-"`
	Owner      int    `json:"owner"`
	Houses     int    `json:"houses"`
	Mortgaged  bool   `json:"mortgaged"`
}

func (property *Property) Landed(game *Game, playerID, _ int) {
	if property.Owner == playerID {
		return
	}

	if property.GetOwned() {
		rent := property.Rent[property.Houses]

		if property.Houses == 0 {
			full := true

			for i := range Sets[property.Set] {
				if game.Board[Sets[property.Set][i]].(*Property).Owner != playerID {
					full = false
					break
				}
			}

			if full {
				rent *= 2
			}
		}

		game.addPayment(playerID, property.Owner, rent)
		return
	}

	game.Queue = append(game.Queue, game.Action)
	game.Action = Action{
		Type: actionPurchaseDecision,
		Data: dataPurchaseDecision{
			PlayerID: playerID,
			Timeout:  time.Now().Add(turnTimeout).Unix(),
		},
	}
}

func (property *Property) GetValue() int {
	return property.Value
}

func (property *Property) GetOwned() bool {
	return property.Owner != NullPlayer
}

func (property *Property) SetOwner(owner int) {
	property.Owner = owner
}

func (property *Property) Purchasable() Purchasable {
	return Purchasable(property)
}

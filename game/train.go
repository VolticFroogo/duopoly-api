package game

import "time"

type Train struct {
	Value     int    `json:"-"`
	Set       int    `json:"-"`
	Rent      [4]int `json:"-"`
	Owner     int    `json:"owner"`
	Mortgaged bool   `json:"mortgaged"`
}

func (train *Train) Landed(game *Game, playerID, _ int) {
	if train.Owner == playerID {
		return
	}

	if train.GetOwned() {
		trainsOwned := 0

		for i := range Sets[train.Set] {
			if game.Board[Sets[train.Set][i]].(*Train).Owner == train.Owner {
				trainsOwned++
			}
		}

		game.addPayment(playerID, train.Owner, train.Rent[trainsOwned-1])
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

func (train *Train) GetValue() int {
	return train.Value
}

func (train *Train) GetOwned() bool {
	return train.Owner != NullPlayer
}

func (train *Train) SetOwner(owner int) {
	train.Owner = owner
}

func (train *Train) Purchasable() Purchasable {
	return Purchasable(train)
}

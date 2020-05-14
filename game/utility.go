package game

import "time"

type Utility struct {
	Value      int   `json:"-"`
	Set        int   `json:"-"`
	Multiplier []int `json:"-"`
	Owner      int   `json:"owner"`
	Mortgaged  bool  `json:"mortgaged"`
}

func (utility *Utility) Landed(game *Game, playerID, roll int) {
	if utility.Owner == playerID {
		return
	}

	if utility.GetOwned() {
		utilitiesOwned := 0

		for i := range Sets[utility.Set] {
			if game.Board[Sets[utility.Set][i]].(*Utility).Owner == utility.Owner {
				utilitiesOwned++
			}
		}

		game.addPayment(playerID, utility.Owner, utility.Multiplier[utilitiesOwned-1]*roll)
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

func (utility *Utility) GetValue() int {
	return utility.Value
}

func (utility *Utility) GetOwned() bool {
	return utility.Owner != NullPlayer
}

func (utility *Utility) SetOwner(owner int) {
	utility.Owner = owner
}

func (utility *Utility) Purchasable() Purchasable {
	return Purchasable(utility)
}

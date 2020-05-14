package game

type Fine struct {
	Value int `json:"-"`
}

func (fine Fine) Landed(game *Game, playerID, _ int) {
	game.addPayment(playerID, NullPlayer, fine.Value)
}

func (fine Fine) Purchasable() Purchasable {
	return nil
}

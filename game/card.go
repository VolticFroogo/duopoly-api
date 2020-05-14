package game

type Card struct {
	Deck int `json:"-"`
}

func (card Card) Landed(game *Game, playerID, _ int) {

}

func (card Card) Purchasable() Purchasable {
	return nil
}

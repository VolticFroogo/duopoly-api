package game

type Jail struct{}

func (jail Jail) Landed(game *Game, playerID, _ int) {

}

func (jail Jail) Purchasable() Purchasable {
	return nil
}

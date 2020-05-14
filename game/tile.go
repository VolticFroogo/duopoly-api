package game

type Tile interface {
	Landed(game *Game, playerID, roll int)
	Purchasable() Purchasable
}

type BlankTile struct{}

func (b BlankTile) Landed(_ *Game, _, _ int) {}
func (b BlankTile) Purchasable() Purchasable {
	return nil
}

package game

type Purchasable interface {
	GetValue() int
	GetOwned() bool
	SetOwner(owner int)
}

package game

import (
	"sync"
)

type Game struct {
	ID          string          `json:"id"`
	OwnerSecret string          `json:"-"`
	Players     map[int]*Player `json:"players"`
	nextPlayer  int             `json:"-"`
	Playing     bool            `json:"playing"`
	mutex       sync.Mutex      `json:"-"`
}

var Games = make(map[string]*Game)

func (game *Game) Run() {
	// Run game timer here.
}

func New(id, secret string) (game *Game) {
	// Create a new game with default values.
	game = &Game{
		ID:          id,
		OwnerSecret: secret,
		Players:     make(map[int]*Player),
		nextPlayer:  0,
		Playing:     false,
		mutex:       sync.Mutex{},
	}

	// Start the game thread on a new goroutine.
	go game.Run()

	// Add the game to the games map.
	Games[id] = game

	return
}

package game

import (
	"math/rand"
	"sync"
	"time"

	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gin-gonic/gin"
)

const (
	NullPlayer    = -1
	tickFrequency = time.Second
	startingMoney = 1500
)

type Game struct {
	ID          string           `json:"id"`
	OwnerSecret string           `json:"-"`
	Players     map[int]*Player  `json:"players"`
	nextPlayer  int              `json:"-"`
	Playing     bool             `json:"playing"`
	Board       *Board           `json:"board,omitempty"`
	Action      Action           `json:"action"`
	Queue       []Action         `json:"-"`
	Deck        [DeckCount][]int `json:"-"`
	mutex       sync.Mutex       `json:"-"`
}

var Games = make(map[string]*Game)

func (game *Game) Run() {
	// Create a ticker.
	t := time.NewTicker(tickFrequency)

	for {
		// Await a tick.
		<-t.C

		game.mutex.Lock()

		if !game.Playing && len(game.Players) == 0 {
			break
		}

		game.Action.Tick(game)

		game.mutex.Unlock()
	}

	// If the game loop has died, we should kill the game.
	// Delete the game from the games map.
	delete(Games, game.ID)
}

func (game *Game) start(playerID int) {
	// Check if the owner is requesting the start.
	if game.OwnerSecret != game.Players[playerID].Secret {
		game.error(playerID, "must be owner to start the game")
		return
	}

	// Set the game to playing.
	game.Playing = true

	// Generate a blank new board for the game.
	game.Board = NewBoard()

	// Initialise all of the players' values to default.
	for i := range game.Players {
		game.Players[i].Position = 0
		game.Players[i].Money = startingMoney
		game.Players[i].Sentence = 0
		game.Players[i].Bankrupt = false
	}

	// Randomly choose the starting player.
	// TODO: implement a dice roll system here or something along those lines
	// which is visible in the opening to the players.
	var startingID int
	startingPlayerPos := rand.Intn(len(game.Players))

	i := 0
	for id := range game.Players {
		if i == startingPlayerPos {
			startingID = id
		}
	}

	game.Action = Action{
		Type: actionTurn,
		Data: dataTurn{
			PlayerID: startingID,
			Timeout:  time.Now().Add(turnTimeout).Unix(),
		},
	}

	// Create the shuffled decks.
	for i := range game.Deck {
		game.Deck[i] = rand.Perm(DeckSize)
	}

	// Broadcast the start of the game.
	game.broadcast(message.Message{
		Type: message.ResponseStart,
		Data: gin.H{
			"startingPlayer": startingID,
		},
	}, nil)
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

	// start the game thread on a new goroutine.
	go game.Run()

	// Add the game to the games map.
	Games[id] = game

	return
}

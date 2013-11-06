package main

import (
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Id              string
	Board           *Board
	Players         []*Player
	Winner          *Player
	CurrentPlayerId string
	playerJoinedCh  chan int
	endTurn         chan int
	mu              sync.Mutex
	isStarted       bool
}

// Creates a new game, but does not start it.
func newGame() *Game {
	id := generateId()

	logInfo("Creating a new game with id: %s.", id)
	game := Game{}

	game.Id = id
	game.Board = &Board{}
	game.isStarted = false

	// init channels
	game.playerJoinedCh = make(chan int)

	board := game.Board
	for i, row := range board.Fields {
		for j, _ := range row {
			board.Fields[i][j] = new(Field)
		}
	}

	return &game
}

func (g *Game) addPlayer(player *Player) {
	g.mu.Lock()
	g.Players = append(g.Players, player)
	g.mu.Unlock()

	if !isTest() && !g.isStarted {
		g.playerJoinedCh <- len(g.Players)
	}
}

// Determines whether a player has already in the game or not.
// Decision is based on the player's name.
func (game *Game) hasAlreadyJoined(player *Player) bool {
	for _, p := range game.Players {
		if p.Name == player.Name {
			return true
		}
	}
	return false
}

// Starts the game.
func (g *Game) start() {
	go g.step()
}

// Returns true if all players are bots in the game.
func (g *Game) isAllBot() bool {
	for _, player := range g.Players {
		if !player.IsBot {
			return false
		}
	}
	return true
}

func (g *Game) step() {
	// wait for enough players to start
	for {
		numberOfPlayers := <-g.playerJoinedCh
		if numberOfPlayers > 1 {
			g.isStarted = true
			break
		}
	}

	// we have enough players so start the first player's turn
	g.doTurn(g.Players[0])

	// main event loop
	for {
		// check if we have a winner
		// when we do, break the main loop
		if result, winner := g.hasWinner(); result {
			g.Winner = winner
			logInfo("The winner is: %s", winner.Name)
			// TODO: handle winner
			break
		}

		prevPlayerId := g.CurrentPlayerId
		player, err := g.getNextPlayer(prevPlayerId)
		if err != nil {
			logError("No player found for id %v", prevPlayerId)
			continue
		}

		g.doTurn(player)
	}

	// TODO: signal the view: game ended, we have a winner
}

// Starts a new turn for a player. Block while the player's turn ends.
func (g *Game) doTurn(player *Player) {
	logInfo("%s started her turn.", player.Name)

	// set current player
	g.CurrentPlayerId = player.Id

	// when the player is a bot, just shoot and end the turn
	if player.IsBot {
		if g.isAllBot() {
			time.Sleep(time.Second * 5)
		}
		g.shootForAI()
	} else {
		// human player have time to react
		g.endTurn = make(chan int)
		go measureTurnTime(g.endTurn)

		<-g.endTurn
	}
	logInfo("%s finished her turn.", player.Name)
}

// Gets the next player in the game. Order based on first came first serve.
func (g *Game) getNextPlayer(playerId string) (*Player, error) {
	for i := range g.Players {
		if g.Players[i].Id == playerId {
			// when we reach the last player, select the first one again
			if len(g.Players) == i+1 {
				return g.Players[0], nil
			} else {
				return g.Players[i+1], nil
			}
		}
	}

	return nil, errorf("No such player: %s", playerId)
}

// Returns true only if there is only one player left on the board. False
// otherwise.
func (g *Game) hasWinner() (bool, *Player) {
	// number of players who have working ships
	x := 0

	var winner *Player

	for _, p := range g.Players {
		if len(p.getWorkingShips()) > 0 {
			x += 1
			winner = p
		}

		if x > 1 {
			return false, nil
		}
	}

	return true, winner
}

func (g *Game) shootForAI() {
	row, col := rand.Intn(SIZE), rand.Intn(SIZE)
	for g.Board.shootAt(row, col, nil) == INVALID {
		row, col = rand.Intn(SIZE), rand.Intn(SIZE)
	}

	logInfo("Bot player shot at %s", rcToS(row, col))
}

func measureTurnTime(endTurn chan int) {
	elapsed := 0.0
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case now := <-ticker.C:
			elapsed = now.Sub(start).Seconds()
			logDebug("Tick: %f", elapsed)
			if elapsed >= TURN_DURATION_SEC {
				ticker.Stop()
				endTurn <- 1
			}
		case <-endTurn:
			return
		}
	}
}

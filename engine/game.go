package engine

import (
	"github.com/kir-dev/torpedo/util"
	"math/rand"
	"sync"
	"time"
)

// Listener interface. Implement this interface, and register views via
// Game#RegisterView to get notifications about certain game events.
type ViewReporter interface {
	// Called when a player after the player shot.
	ReportHitResult(row, col int, result HitResult)
	// Called when there are enough players to start the game.
	ReportGameStarted()
	// Called when the game ends.
	ReportGameOver(winner *Player)
	// Called on every second with the already elapsed seconds in the player's
	// turn.
	ReportElapsedTime(elapsed float64)
	// Called when a new turn starts. Reports the current and the next player.
	ReportPlayerTurnStart(current *Player, next *Player)
	// Called when a new player joins the game.
	ReportPlayerJoined(player *Player)
}

type Game struct {
	Id              string
	Board           *Board
	Players         []*Player
	Winner          *Player
	CurrentPlayerId string

	views          []ViewReporter
	endCh          chan<- int
	playerJoinedCh chan int
	endTurn        chan int
	isStarted      bool

	mu     sync.Mutex
	viewMu sync.Mutex
}

// Creates a new game with a channel on which it will signal when the game ends.
func NewGame(end chan<- int) *Game {
	g := newGame()
	g.endCh = end

	return g
}

// Starts the game.
func (g *Game) Start() {
	go g.step()
}

func (g *Game) Shoot(row, col int) HitResult {
	result := g.Board.shootAt(row, col, g.endTurn)
	g.notifyViewsAfterShot(row, col, result)
	return result
}

// registers a view for the game
func (g *Game) RegisterView(reporter ViewReporter) {
	g.viewMu.Lock()
	defer g.viewMu.Unlock()

	g.views = append(g.views, reporter)
}

func (g *Game) DiscardView(view ViewReporter) {
	g.viewMu.Lock()
	defer g.viewMu.Unlock()

	var index int

	// find the view in question
	for i, v := range g.views {
		if v == view {
			index = i
		}
	}

	// delete it
	g.views = append(g.views[:index], g.views[index+1:]...)
}

func (g *Game) GetPlayerById(id string) *Player {
	for _, player := range g.Players {
		if player.Id == id {
			return player
		}
	}

	return nil
}

func (g *Game) RemovePlayer(id string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	var index int
	index = -1

	// find the player in question
	for i, v := range g.Players {
		if v.Name == id {
			index = i
		}
	}

	// delete it
	if(index > -1) {
		g.Players = append(g.Players[:index], g.Players[index+1:]...)
	}
}

// Creates a new game, but does not start it.
func newGame() *Game {
	id := generateId()

	util.LogInfo("Creating a new game with id: %s.", id)
	game := Game{}

	game.Id = id
	fields := make([][]*Field, conf.BoardSize)
	for i := range fields {
		fields[i] = make([]*Field, conf.BoardSize)
	}
	game.Board = &Board{Fields: fields}
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

	if !util.IsTest() && !g.isStarted {
		g.playerJoinedCh <- len(g.Players)
	}

	for _, view := range g.views {
		view.ReportPlayerJoined(player)
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
		if numberOfPlayers > conf.MinimalPlayerCnt-1 {
			g.isStarted = true
			break
		}
	}

	// notify views that the game started
	for _, reporter := range g.views {
		reporter.ReportGameStarted()
	}

	// we have enough players so start the first player's turn
	g.reportTurnStart(g.Players[0], g.Players[1])
	g.doTurn(g.Players[0])

	// game event loop
	for {
		// check if we have a winner
		// when we do, break the game loop
		if result, winner := g.hasWinner(); result {
			g.Winner = winner
			util.LogInfo("The winner is: %s", winner.Name)
			break
		}

		prevPlayerId := g.CurrentPlayerId
		player, err := g.getNextPlayer(prevPlayerId)
		if err != nil {
			util.LogError("No player found for id %v", prevPlayerId)
			continue
		}

		// report new turn for player
		nextPlayer, _ := g.getNextPlayer(player.Id)
		g.reportTurnStart(player, nextPlayer)

		g.doTurn(player)
	}

	// notify views
	for _, reporter := range g.views {
		reporter.ReportGameOver(g.Winner)
	}
	g.endCh <- 1 // signal the main loop
	g.cleanUp()
}

// Starts a new turn for a player. Block while the player's turn ends.
func (g *Game) doTurn(player *Player) {
	util.LogInfo("%s started her turn.", player.Name)

	// set current player
	g.CurrentPlayerId = player.Id

	// when the player is a bot, just shoot and end the turn
	if player.IsBot {
		if g.isAllBot() && conf.WaitForBots {
			time.Sleep(time.Second * time.Duration(conf.BotTurnDurationSec))
		}
		g.shootForAI()
	} else {
		// human player have time to react
		g.endTurn = make(chan int)
		go g.measureTurnTime()

		<-g.endTurn
	}
	util.LogInfo("%s finished her turn.", player.Name)
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

	return nil, util.Errorf("No such player: %s", playerId)
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
	row, col := rand.Intn(conf.BoardSize), rand.Intn(conf.BoardSize)
	result := g.Board.shootAt(row, col, nil)

	for result == INVALID {
		row, col = rand.Intn(conf.BoardSize), rand.Intn(conf.BoardSize)
		result = g.Board.shootAt(row, col, nil)
	}

	g.notifyViewsAfterShot(row, col, result)

	util.LogInfo("Bot player shot at %s and the result is: %s", RowColToS(row, col), result)
}

func (g *Game) measureTurnTime() {
	elapsed := 0.0
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case now := <-ticker.C:
			elapsed = now.Sub(start).Seconds()
			util.LogDebug("Tick: %f", elapsed)

			// report tick
			for _, reporter := range g.views {
				reporter.ReportElapsedTime(elapsed)
			}

			if elapsed >= float64(conf.TurnDurationSec) {
				ticker.Stop()
				g.endTurn <- 1
			}
		case <-g.endTurn:
			return
		}
	}
}

// notfiy the the views after a shot has been fired
func (g *Game) notifyViewsAfterShot(row, col int, result HitResult) {
	for _, reporter := range g.views {
		reporter.ReportHitResult(row, col, result)
	}
}

func (g *Game) reportTurnStart(current, next *Player) {
	for _, view := range g.views {
		view.ReportPlayerTurnStart(current, next)
	}
}

func (g *Game) cleanUp() {
	// release views
	g.views = nil

	// close channels
	close(g.playerJoinedCh)
	// NOTE: g.endTurn was closed at the end of the last turn

	ResetColors()
}

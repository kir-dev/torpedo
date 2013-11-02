package main

type Game struct {
	Board           Board
	Players         []*Player
	CurrentPlayerId string
}

var (
	currentGame Game
)

type shootResult string

// Starts a new game.
func startNewGame() Game {
	logInfo("Starting a new game.")
	currentGame = Game{}

	board := &currentGame.Board

	for i, row := range board.Fields {
		for j, _ := range row {
			board.Fields[i][j] = new(Field)
		}
	}

	return currentGame
}

func (g *Game) addPlayer(player *Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) shootAt(row, col int) shootResult {
	field := g.Board.Fields[row][col]

	if field.IsHit {
		return shootResult("invalid")
	}

	result := shootResult("miss")
	field.IsHit = true
	if !field.IsEmpty() {
		field.ShipPart.IsHit = true
		result = shootResult("hit")
	}

	return result
}

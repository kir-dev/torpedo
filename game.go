package main

type Game struct {
	Board   Board
	Players []*Player
}

var (
	currentGame Game
)

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

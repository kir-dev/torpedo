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
	currentGame = Game{}

	board := &currentGame.Board

	for i, row := range currentGame.Board.Fields {
		for j, _ := range row {
			board.Fields[i][j] = new(Field)
		}
	}

	return currentGame
}

func (g *Game) addPlayer(player *Player) {
	g.Players = append(g.Players, player)
}

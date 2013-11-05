package main

type Game struct {
	Board           Board
	Players         []*Player
	CurrentPlayerId string
}

type hitResult string

// Creates a new game, but does not start it.
func newGame() *Game {
	logInfo("Creating a new game.")
	game := Game{}

	board := &game.Board

	for i, row := range board.Fields {
		for j, _ := range row {
			board.Fields[i][j] = new(Field)
		}
	}

	return &game
}

func (g *Game) addPlayer(player *Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) shootAt(row, col int) hitResult {
	field := g.Board.Fields[row][col]

	if field.IsHit {
		return hitResult("invalid")
	}

	result := hitResult("miss")
	field.IsHit = true
	if !field.IsEmpty() {
		field.ShipPart.IsHit = true
		result = hitResult("hit")
	}

	return result
}

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

}

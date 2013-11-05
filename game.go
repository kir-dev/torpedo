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

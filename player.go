package main

type Player struct {
	Name  string
	IsBot bool
	Ships []*Ship
	Id    string
}

// creates a new player
func newPlayer(name string) *Player {
	// TODO check for ID uniqueness
	return &Player{name, false, nil, generateId()}
}

// Adds a player to the registry.
func (player *Player) join(game *Game) error {
	logInfo("New player (with name: %s) attempts to join the game.", player.Name)
	if game.hasAlreadyJoined(player) {
		return errorf("Player with name %s has already joined the game.", player.Name)
	}

	err := game.Board.placeShips(game.Players, player)
	if err != nil {
		return err
	}

	game.addPlayer(player)

	logInfo("Player with name %s has joined the game.", player.Name)
	return nil
}

func (p Player) getWorkingShips() []*Ship {
	ships := make([]*Ship, 0, len(p.Ships))
	for _, ship := range p.Ships {
		if !ship.isSunken() {
			ships = append(ships, ship)
		}
	}

	return ships
}

func (player Player) getCurrentScore() float64 {
	var sum float64 = 0
	for _, ship := range player.getWorkingShips() {
		sum += ship.getScore()
	}
	return sum
}

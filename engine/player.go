package engine

import (
	"github.com/kir-dev/torpedo/util"
)

type Player struct {
	Name  string
	IsBot bool
	Ships []*Ship
	Id    string
	Color Color
}

// creates a new player
func NewPlayer(name string) *Player {
	return &Player{name, false, nil, generateId(), getNextColor()}
}

// Adds a player to the registry.
func (player *Player) Join(game *Game) error {
	util.LogInfo("New player (with name: %s) attempts to join the game.", player.Name)
	if game.hasAlreadyJoined(player) {
		return util.Errorf("Player with name %s has already joined the game.", player.Name)
	}

	err := game.Board.placeShips(game.Players, player)
	if err != nil {
		return err
	}

	game.addPlayer(player)

	util.LogInfo("Player with name %s (id: %s) has joined the game.", player.Name, player.Id)
	return nil
}

func (p Player) getWorkingShips() []*Ship {
	ships := make([]*Ship, 0, len(p.Ships))
	for _, ship := range p.Ships {
		if !ship.IsSunken() {
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

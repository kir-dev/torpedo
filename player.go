package main

import (
	"log"
)

type Player struct {
	Name  string
	IsBot bool
	Ships []*Ship
}

// creates a new player
func newPlayer(name string) *Player {
	return &Player{name, false, nil}
}

// Adds a player to the registry.
func join(player *Player) error {
	if player.hasAlreadyJoined() {
		return errorf("Player with name %s has already joined the game.", player.Name)
	}

	currentGame.Board.placeShips(player)
	currentGame.addPlayer(player)

	log.Printf("Player with name %s has joined the game.", player.Name)

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

func (player *Player) hasAlreadyJoined() bool {
	for _, p := range currentGame.Players {
		if p.Name == player.Name {
			return true
		}
	}
	return false

}

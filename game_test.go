package main

import (
	"testing"
)

func TestAddingPlayerToGame(t *testing.T) {
	g := Game{}

	playerCount := len(g.Players)

	g.addPlayer(newPlayer("Teszt Elek"))

	if playerCount >= len(g.Players) {
		t.Error("Player was not added to the game")
	}
}

func TestInitBoardOnStart(t *testing.T) {
	g := startNewGame()

	for i, row := range g.Board.Fields {
		for j, f := range row {
			if f == nil {
				t.Errorf("%s:%s field was not initialized", i, j)
			}
		}
	}
}

package main

import (
	"testing"
)

func TestAddingPlayerToGame(t *testing.T) {
	g := newGame()
	playerCount := len(g.Players)

	g.addPlayer(newPlayer("Teszt Elek"))

	if playerCount >= len(g.Players) {
		t.Error("Player was not added to the game")
	}
}

func TestInitBoardOnStart(t *testing.T) {
	g := newGame()

	for i, row := range g.Board.Fields {
		for j, f := range row {
			if f == nil {
				t.Errorf("%s:%s field was not initialized", i, j)
			}
		}
	}
}

func TestGetNextPlayer(t *testing.T) {
	g := newGame()

	player1 := newPlayer("t1")
	player2 := newPlayer("t2")
	player3 := newPlayer("t3")

	player1.join(g)
	player2.join(g)
	player3.join(g)

	p, _ := g.getNextPlayer(player2.Id)

	if p != player3 {
		t.Errorf("Next player should be %v, but got %v.", player3, p)
	}
}

func TestGetNextPlayerReturnsErrorWhenNoPlayerFound(t *testing.T) {
	g := newGame()

	_, err := g.getNextPlayer("test")

	if err == nil {
		t.Error("Should have returned an error when player is not found.")
	}
}

func TestGetNextPlayerShouldReturnFirstPlayerIfCurrentIsTheLast(t *testing.T) {
	g := newGame()

	player1 := newPlayer("t1")
	player2 := newPlayer("t2")

	player1.join(g)
	player2.join(g)

	p, _ := g.getNextPlayer(player2.Id)

	if p != player1 {
		t.Errorf("Next player should be %v, but got %v.", player1, p)
	}
}

func TestIsAllBotWithOneHumanPlayer(t *testing.T) {
	g := newGame()

	p := newPlayer("t1")
	p.IsBot = true
	p.join(g)

	newPlayer("t2").join(g)

	if g.isAllBot() == true {
		t.Errorf("There is one human player, should've returned false.")
	}
}

func TestIsAllBot(t *testing.T) {
	g := newGame()

	p := newPlayer("t1")
	p.IsBot = true
	p.join(g)

	p = newPlayer("t2")
	p.IsBot = true
	p.join(g)

	if g.isAllBot() == false {
		t.Errorf("There is no human player, should've returned true.")
	}
}

func TestHasWinner(t *testing.T) {
	g := newGame()
	p1 := newPlayer("t1")
	p1.Ships = append(p1.Ships, newShip(2))
	p2 := newPlayer("t2")

	g.Players = append(g.Players, p1, p2)

	result, winner := g.hasWinner()

	if !result {
		t.Fatal("No winner found, but there should be one.")
	}
	if winner != p1 {
		t.Errorf("Expected winner: %s, got: %s", p1.Name, winner.Name)
	}
}

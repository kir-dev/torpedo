package engine

import (
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestPlayerJoins(t *testing.T) {
	game := newGame()
	NewPlayer("test").Join(game)

	if len(game.Players) < 1 {
		t.Error("Player could not join the game.")
	}
}

func TestPlayerCannotJoinTwice(t *testing.T) {
	game := newGame()
	player := NewPlayer("test")

	player.Join(game)
	player.Join(game)

	if len(game.Players) > 1 {
		t.Error("Player joined twice.")
	}
}

func TestPlayerHasAlredyJoined(t *testing.T) {
	game := newGame()

	player := NewPlayer("test")
	player.Join(game)

	if !game.hasAlreadyJoined(player) {
		t.Error("Player should have been marked as joined.")
	}
}

func TestPlayerHasAlreadyJoinedIsIndependentFromThePlayerObject(t *testing.T) {
	game := newGame()

	NewPlayer("test").Join(game)
	player := NewPlayer("test")

	if !game.hasAlreadyJoined(player) {
		t.Error("Player should have been marked as joined.")
	}
}

func TestGetTheNumberOfWorkingShips(t *testing.T) {
	p := NewPlayer("t1")
	if n := len(p.getWorkingShips()); n != 0 {
		t.Errorf("Expected 0 but got %d", n)
	}
}

func TestGetWorkingShips(t *testing.T) {
	p := NewPlayer("t1")
	p.Ships = append(p.Ships, newShip(2))

	if n := len(p.getWorkingShips()); n != 1 {
		t.Errorf("Expected 1 but got %d", n)
	}

}

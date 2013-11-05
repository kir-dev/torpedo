package main

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
	newPlayer("test").join(game)

	if len(game.Players) < 1 {
		t.Error("Player could not join the game.")
	}
}

func TestPlayerCannotJoinTwice(t *testing.T) {
	game := newGame()
	player := newPlayer("test")

	player.join(game)
	player.join(game)

	if len(game.Players) > 1 {
		t.Error("Player joined twice.")
	}
}

func TestPlayerHasAlredyJoined(t *testing.T) {
	game := newGame()

	player := newPlayer("test")
	player.join(game)

	if !game.hasAlreadyJoined(player) {
		t.Error("Player should have been marked as joined.")
	}
}

func TestPlayerHasAlreadyJoinedIsIndependentFromThePlayerObject(t *testing.T) {
	game := newGame()

	newPlayer("test").join(game)
	player := newPlayer("test")

	if !game.hasAlreadyJoined(player) {
		t.Error("Player should have been marked as joined.")
	}
}

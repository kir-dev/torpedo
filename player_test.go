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
	startNewGame()
	join(newPlayer("test"))

	if len(currentGame.Players) < 1 {
		t.Error("Player could not join the game.")
	}
}

func TestPlayerCannotJoinTwice(t *testing.T) {
	startNewGame()
	player := newPlayer("test")

	join(player)
	join(player)

	if len(currentGame.Players) > 1 {
		t.Error("Player joined twice.")
	}
}

func TestPlayerHasAlredyJoined(t *testing.T) {
	startNewGame()

	player := newPlayer("test")
	join(player)

	if !player.hasAlreadyJoined() {
		t.Error("Player should have been marked as joined.")
	}
}

func TestPlayerHasAlreadyJoinedIsIndependentFromThePlayerObject(t *testing.T) {
	startNewGame()

	join(newPlayer("test"))
	player := newPlayer("test")

	if !player.hasAlreadyJoined() {
		t.Error("Player should have been marked as joined.")
	}
}

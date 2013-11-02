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
	StartNewGame()
	Join(NewPlayer("test"))

	if len(currentGame.Players) < 1 {
		t.Error("Player could not join the game.")
	}
}

func TestPlayerCannotJoinTwice(t *testing.T) {
	StartNewGame()
	player := NewPlayer("test")

	Join(player)
	Join(player)

	if len(currentGame.Players) > 1 {
		t.Error("Player joined twice.")
	}
}

func TestPlayerHasAlredyJoined(t *testing.T) {
	StartNewGame()

	player := NewPlayer("test")
	Join(player)

	if !player.hasAlreadyJoined() {
		t.Error("Player should have been marked as joined.")
	}
}

func TestPlayerHasAlreadyJoinedIsIndependentFromThePlayerObject(t *testing.T) {
	StartNewGame()

	Join(NewPlayer("test"))
	player := NewPlayer("test")

	if !player.hasAlreadyJoined() {
		t.Error("Player should have been marked as joined.")
	}
}

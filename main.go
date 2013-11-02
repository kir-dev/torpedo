package main

import (
	"log"
	"time"
	// "net/http"
	"math/rand"
	"os"
)

const (
	ENV = "ENV"
	DEV = "development"
)

type content struct {
	Title string
}

func main() {
	rand.Seed(time.Now().Unix())
	log.SetOutput(os.Stdout)
	// fmt.Println("Starting on port 8080...")

	// log.Fatal(http.ListenAndServe(":8080", nil))
	startNewGame()
	join(newPlayer("Player1"))
	join(newPlayer("Player2"))
	join(newPlayer("Player3"))
	join(newPlayer("Player4"))
	currentGame.Board.print()

}

func isDev() bool {
	return os.Getenv(ENV) == "" || os.Getenv(ENV) == DEV
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	ENV = "ENV"
	DEV = "development"
)

type content struct {
	Title string
}

func main() {
	fmt.Println("Starting on port 8080...")
	fmt.Println("Press Ctrl-C to exit!")

	rand.Seed(time.Now().Unix())
	log.SetOutput(os.Stdout)
	startNewGame()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isDev() bool {
	return os.Getenv(ENV) == "" || os.Getenv(ENV) == DEV
}

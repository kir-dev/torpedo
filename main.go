package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	ENV  = "ENV"
	DEV  = "development"
	TEST = "test"
)

var (
	currentGame *Game

	configPath = flag.String("config", "", "Path of the config file. If left empty, default config will be loaded.")
	port       = flag.String("port", ":8080", "Port to bind to.")
	conf       = defaultConfig()
)

func main() {
	rand.Seed(time.Now().Unix())
	log.SetOutput(os.Stdout)

	flag.Parse()

	fmt.Printf("Starting on port %s...\n", *port)
	if isDev() {
		fmt.Println("Started in DEVELOPMENT mode.")
	}
	fmt.Println("Press Ctrl-C to exit!")
	fmt.Println()

	// load config
	if *configPath != "" {
		conf = loadConfig(*configPath)
	}

	logInfo("Loaded config: %#v", conf)

	// start main loop for the application
	go mainLoop()

	http.Handle("/public/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(*port, nil))

}

// main loop for the program, starts new game when needed
func mainLoop() {
	endGame := make(chan int)

	for {
		currentGame = newGameWithEndChannel(endGame)
		currentGame.start()

		<-endGame
		// TODO: archive game before creating a new one
	}
}

func isDev() bool {
	return os.Getenv(ENV) == "" || os.Getenv(ENV) == DEV
}

func isTest() bool {
	return os.Getenv(ENV) == TEST
}

package main

import (
	"flag"
	"fmt"
	"github.com/kir-dev/torpedo/engine"
	"github.com/kir-dev/torpedo/util"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	currentGame *engine.Game   = nil
	port                       = flag.String("port", ":8080", "Port to bind to.")
	Archive     []*engine.Game = nil
)

func main() {
	fmt.Printf("Starting on port %s...\n", *port)
	if util.IsDev() {
		fmt.Println("Started in DEVELOPMENT mode.")
	}
	fmt.Println("Press Ctrl-C to exit!")
	fmt.Println()

	initialize()

	// start main loop for the application
	go mainLoop()

	http.Handle("/public/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(*port, nil))
}

func initialize() {
	// seed with current time.
	rand.Seed(time.Now().Unix())
	// log to stdout instead of stderr
	log.SetOutput(os.Stdout)

	// load config
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path of the config file. If left empty, default config will be loaded.")
	flag.Parse()
	engine.LoadConfig(configPath)
}

// main loop for the program, starts new game when needed
func mainLoop() {
	endGame := make(chan int)

	for {
		currentGame = engine.NewGame(endGame)
		currentGame.Start()

		<-endGame
		Archive = append(Archive, currentGame)
		util.LogInfo("Game ended with id: %s.", currentGame.Id)
	}
}

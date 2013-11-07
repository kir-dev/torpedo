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
	configPath  = flag.String("config", "", "Path of the config file. If left empty, default config will be loaded.")
	port        = flag.String("port", ":8080", "Port to bind to.")
	conf        = defaultConfig()
)

type content struct {
	Title string
}

func main() {
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

	rand.Seed(time.Now().Unix())
	log.SetOutput(os.Stdout)
	currentGame = newGame()
	currentGame.start()

	http.Handle("/public/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(*port, nil))

}

func isDev() bool {
	return os.Getenv(ENV) == "" || os.Getenv(ENV) == DEV
}

func isTest() bool {
	return os.Getenv(ENV) == TEST
}

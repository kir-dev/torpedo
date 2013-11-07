package main

import (
	"encoding/json"
	"os"
)

const (
	TURN_DURATION_SEC = 30
	BOT_TURN_DURATION = 5
)

type config struct {
	// Maximum duration of a turn for a player. This does not apply to bots.
	TurnDurationSec int `json:"turn_duration"`

	// Duration for all bot game turns.
	BotTurnDurationSec int `json:"bot_turn_duration"`

	// In all bot games we stall the game for every bot player if this flag is
	// set to true.
	WaitForBots bool `json:"wait_for_bots"`
}

func loadConfig(path string) config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	rawContent := make([]byte, stat.Size())
	_, err = f.Read(rawContent)

	if err != nil {
		panic(err)
	}

	return loadConfigFromBytes(rawContent)
}

func loadConfigFromBytes(rawContent []byte) config {
	conf = defaultConfig()

	err := json.Unmarshal(rawContent, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func defaultConfig() config {
	return config{
		TurnDurationSec:    TURN_DURATION_SEC,
		BotTurnDurationSec: BOT_TURN_DURATION,
		WaitForBots:        true,
	}
}

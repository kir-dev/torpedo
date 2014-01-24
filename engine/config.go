package engine

import (
	"encoding/json"
	"github.com/kir-dev/torpedo/util"
	"os"
)

const (
	TURN_DURATION_SEC  = 30
	BOT_TURN_DURATION  = 5
	BOARD_SIZE         = 26
	MINIMAL_PLAYER_CNT = 2
)

var (
	conf config
)

type config struct {
	// Maximum duration of a turn for a player. This does not apply to bots.
	TurnDurationSec int `json:"turn_duration"`

	// Duration for all bot game turns.
	BotTurnDurationSec int `json:"bot_turn_duration"`

	// In all bot games we stall the game for every bot player if this flag is
	// set to true.
	WaitForBots bool `json:"wait_for_bots"`

	// default board size
	BoardSize int `json:"board_size"`

	// minimal players count to start game
	MinimalPlayerCnt int `json: "minimal_player_cnt"`
}

func LoadConfig(path string) {
	if path == "" {
		util.LogDebug("Loaded default config.")
		conf = defaultConfig()
		return
	}

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

	conf = loadConfigFromBytes(rawContent)
}

func loadConfigFromBytes(rawContent []byte) config {
	conf = defaultConfig()

	err := json.Unmarshal(rawContent, &conf)
	if err != nil {
		panic(err)
	}

	util.LogDebug("Loaded config: %#v", conf)

	return conf
}

func defaultConfig() config {
	return config{
		TurnDurationSec:    TURN_DURATION_SEC,
		BotTurnDurationSec: BOT_TURN_DURATION,
		WaitForBots:        true,
		BoardSize:          BOARD_SIZE,
		MinimalPlayerCnt:   MINIMAL_PLAYER_CNT,
	}
}

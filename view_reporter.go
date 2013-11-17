package main

import (
	"github.com/gorilla/websocket"
	"github.com/kir-dev/torpedo/engine"
	"github.com/kir-dev/torpedo/util"
)

type messageType int

const (
	MSG_HITRESULT messageType = iota
	MSG_GAMESTARTED
	MSG_GAMEOVER
	MSG_ELAPSEDTIME
	MSG_TURNSTART
	MSG_PLAYERJOINED
)

const MAX_VIEW_ERROR_COUNT = 3

type viewReporter struct {
	conn       *websocket.Conn
	remoteAddr string
	errorCount int
}

type viewMsg struct {
	Type    messageType `json:"type"`
	Payload interface{} `json:"payload"`
}

func (v *viewReporter) ReportHitResult(row, col int, result engine.HitResult) {
	values := map[string]interface{}{
		"row":    row,
		"col":    col,
		"result": result,
	}

	v.send(MSG_HITRESULT, values)
}

func (v *viewReporter) ReportGameStarted() {
	v.send(MSG_GAMESTARTED, nil)
}

func (v *viewReporter) ReportGameOver(winner *engine.Player) {
	v.send(MSG_GAMEOVER, winner.Name)
}

func (v *viewReporter) ReportElapsedTime(elapsed float64) {
	v.send(MSG_ELAPSEDTIME, elapsed)
}

func (v *viewReporter) ReportPlayerJoined(player *engine.Player) {
	v.send(MSG_PLAYERJOINED, player.Name)
}

func (v *viewReporter) ReportPlayerTurnStart(current *engine.Player, next *engine.Player) {
	var nextPlayerName string
	if next == nil {
		nextPlayerName = "error: could not determine next player"
	} else {
		nextPlayerName = next.Name
	}

	names := map[string]string{
		"current": current.Name,
		"next":    nextPlayerName,
	}

	v.send(MSG_TURNSTART, names)
}

func (v *viewReporter) send(aType messageType, payload interface{}) {
	// try to write in the socket
	err := v.conn.WriteJSON(viewMsg{aType, payload})

	if err != nil {
		if v.errorCount >= MAX_VIEW_ERROR_COUNT {
			currentGame.DiscardView(v)
			util.LogWarn("Discarded view for %s after %d errors", v.remoteAddr, v.errorCount)
			return
		}

		util.LogWarn("Could not send data to %s. Error msg: %s", v.remoteAddr, err)
		v.errorCount += 1
	}
}

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
)

type viewReporter struct {
	conn       *websocket.Conn
	remoteAddr string
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

	if err := v.conn.WriteJSON(viewMsg{MSG_HITRESULT, values}); err != nil {
		// TODO: discard view after n errors
		util.LogError("Could not send data to %s.", v.remoteAddr)
	}
}

func (v *viewReporter) ReportGameStarted() {

}

func (v *viewReporter) ReportGameOver(winner *engine.Player) {

}

func (v *viewReporter) ReportElapsedTime(elapsed float64) {

}

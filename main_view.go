package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

const (
	HEADER_ORIGIN = "Origin"
	READ_BUFFER   = 1024
	WRITE_BUFFER  = 1024
)

func init() {
	http.HandleFunc("/ws", websocketHandler)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// enforce same origin policy
	if r.Header.Get(HEADER_ORIGIN) != "http://"+r.Host {
		http.Error(w, "Origin not allowed", http.StatusForbidden)
		return
	}

	// create websocket
	conn, err := websocket.Upgrade(w, r, w.Header(), READ_BUFFER, WRITE_BUFFER)
	if err != nil {
		http.Error(w, "Could not upgrade the connection.", http.StatusBadRequest)
		return
	}

	currentGame.RegisterView(&viewReporter{conn, r.RemoteAddr, 0})
}

package main

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type Command struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func eventHandler(ws *websocket.Conn) {

	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	event := msg[:n]

	ws.Write(msg)
}

func main() {
	http.Handle("/", websocket.Handler(eventHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

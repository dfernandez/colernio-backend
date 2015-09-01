package server

import (
	"encoding/json"
	"log"
	"strconv"

	"golang.org/x/net/websocket"
)

type Command struct {
	Event string `json:"event"`
	Data  string `json:"data,omitempty"`
}

type Server struct {
	Clients map[string]Client
}

var Srv Server

func (srv Server) broadcast(cmd Command) {
	for _, c := range srv.Clients {
		c.write(cmd)
	}
}

func (srv Server) register(c Client) {
	srv.Clients[c.key] = c
	srv.broadcast(Command{"num_clients", strconv.Itoa(len(srv.Clients))})
}

func (srv Server) unregister(c Client) {
	delete(srv.Clients, c.key)
	srv.broadcast(Command{"num_clients", strconv.Itoa(len(srv.Clients))})
}

func WebsocketHandler(ws *websocket.Conn) {
	key := ws.Request().Header.Get("Sec-Websocket-Key")

	client := Client{ws, key}
	Srv.register(client)

	defer func() {
		Srv.unregister(client)
	}()

	for {
		msg := make([]byte, 512)
		n, err := ws.Read(msg)

		if err != nil {
			break
		}

		var cmd Command
		json.Unmarshal(msg[:n], &cmd)

		log.Println(cmd.Event)

		client.process(cmd)
	}
}

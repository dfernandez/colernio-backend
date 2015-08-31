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

type Client struct {
	ws  *websocket.Conn
	key string
}

type Server struct {
	Clients map[string]Client
}

var Srv Server

func (srv Server) broadcast(cmd Command) {
	response, _ := json.Marshal(cmd)

	for _, c := range srv.Clients {
		c.write(response)
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

func (c Client) process(cmd Command) {
	log.Println(cmd.Event)

	response, _ := json.Marshal(Command{"pong", ""})
	c.write(response)
}

func (c Client) write(response []byte) {
	c.ws.Write(response)
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

		client.process(cmd)
	}
}

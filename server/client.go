package server

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

type Client struct {
	ws  *websocket.Conn
	key string
}

func (c Client) process(cmd Command) {
	switch cmd.Event {
	case "ping":
		c.write(Command{"pong", ""})
	}
}

func (c Client) write(cmd Command) {
	response, _ := json.Marshal(cmd)
	c.ws.Write(response)
}

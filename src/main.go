package main

import (
	"github.com/IggySamma/Web-IRC/ws"
)

func main() {
	ws.StartServer(ws.MessageHandler)
	channels := &ws.Channel{channel: make(map[string]*ws.LinkedList)}

	for {
	}
}

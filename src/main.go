package main

import (
	"github.com/IggySamma/Web-IRC/ws"
)

func main() {
	ws.StartServer(ws.MessageHandler)
	for {
	}
}

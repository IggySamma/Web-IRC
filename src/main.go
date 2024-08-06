package main

import (
	"log"

	"github.com/IggySamma/Web-IRC/ws"
)

func main() {
	server := ws.StartServer(ws.MessageHandler)

	if server == nil {
		log.Fatal("Failed to start the server")
		return
	}

	log.Println("Server started...")

	// Block the main function indefinitely
	select {}
}

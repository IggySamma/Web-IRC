package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var websocketBuffer = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", websocketHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func websocketHandler(writer http.ResponseWriter, reader *http.Request) {
	connection, err := websocketBuffer.Upgrade(writer, reader, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Received message: ", string(message))

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

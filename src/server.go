package main

import (
	"log"
	"net/http"
)

func main() {
	path := http.FileServer(http.Dir("./client"))
	http.Handle("/", path)
	http.HandleFunc("/ws", websocketHandler)
	log.Println(http.ListenAndServe(":3000", nil))
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := WebsocketBuffer.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer connection.Close()

	if !CheckForUsername(connection.RemoteAddr()) {
		messageType, _, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		err = connection.WriteMessage(messageType, []byte("Please enter your username: "))
		if err != nil {
			log.Println(err)
			return
		}
	}

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		parsed := []byte(MessageHandler(connection.RemoteAddr(), string(message)))
		err = connection.WriteMessage(messageType, parsed)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

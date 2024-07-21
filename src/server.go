package main

import (
	"fmt"
	"log"
	"net/http"

	/*"strconv"*/

	/*"os"*/
	"strings"

	"github.com/gorilla/websocket"
)

var websocketBuffer = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//Temp for test remove later
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	path := http.FileServer(http.Dir("./client"))
	http.Handle("/", path)
	http.HandleFunc("/ws", websocketHandler)
	log.Println(http.ListenAndServe(":3000", nil))
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := websocketBuffer.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(connection.RemoteAddr())

	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(connection.LocalAddr())
		fmt.Println("Received message: ", string(message))
		parsed := messageHandler(string(message))
		fmt.Println("Parsed message: ", string(messageHandler(string(message))))
		fmt.Println("Parsed message ? : ", parsed)

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func messageHandler(message string) string {
	if message != "" {
		log.Println(message)
	}
	if strings.Contains(message, "Username: ") {
		return addUser(message)
	}
	return message
}

func addUser(username string) string {
	username = strings.TrimLeft(username, "Username: ")
	trimmed, found := strings.CutSuffix(username, " ")
	if !found && len(trimmed) == 0 {
		log.Println(trimmed)
		return "Username is blank"
	}
	return trimmed
}

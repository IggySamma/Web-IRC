package main

import (
	"fmt"
	"log"
	"net"
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

var usernames = make(map[net.Addr]string)

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
	/*if !checkForUsername(connection.LocalAddr()) {
		return "Enter username: "
	}*/
	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Received message: ", string(message))
		parsed := []byte(messageHandler(connection.LocalAddr(), string(message)))
		fmt.Println("Parsed message : ", parsed)

		err = connection.WriteMessage(messageType, parsed)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func messageHandler(connection net.Addr, message string) string {
	if message != "" {
		log.Println("Failed: " + message)
	}
	if !checkForUsername(connection) {
		return "Enter username: "
	}
	if strings.Contains(message, "Username: ") {
		return setUsername(connection, message)
	}
	return message
}

func setUsername(connection net.Addr, message string) string {
	username := usernames[connection]
	username = strings.TrimLeft(message, "Username: ")
	username, found := strings.CutSuffix(username, " ")
	if !found && len(username) == 0 {
		log.Println(username)
		return "Username is blank, please try again"
	}
	usernames[connection] = username

	return "Username set as: " + username
}

func checkForUsername(connection net.Addr) bool {
	_, check := usernames[connection]
	return check
}

/*
func sendMessage()*/

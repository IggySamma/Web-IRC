package ws

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Server struct {
	client        map[*websocket.Conn]string
	handleMessage func(server *Server, connection *websocket.Conn, message []byte)
}

var WebsocketBuffer = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//Temp for test remove later
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer(handleMessage func(server *Server, connection *websocket.Conn, message []byte)) *Server {
	path := http.FileServer(http.Dir("./client"))

	http.Handle("/", path)

	server := Server{
		make(map[*websocket.Conn]string),
		handleMessage,
	}

	http.HandleFunc("/ws", server.WebsocketHandler)

	go http.ListenAndServe(":3000", nil)

	return &server
}

func MessageHandler(server *Server, connection *websocket.Conn, message []byte) {
	if !server.CheckForUsername(connection) {
		server.Reply([]byte("Enter username: "))
	}

	if strings.Contains(string(message), "Username: ") {
		server.Reply([]byte(server.SetUsername(connection, message)))
	}

	if server.CheckForUsername(connection) && string(message) != "" {
		log.Println(string(message))
		server.Reply(message)
	}
}

func (server *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := WebsocketBuffer.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	defer connection.Close()

	server.client[connection] = ""

	for {
		messageType, message, err := connection.ReadMessage()

		if err != nil || messageType == websocket.CloseMessage {
			break
		}

		go server.handleMessage(server, connection, message)
	}

	delete(server.client, connection)
}

func (server *Server) Reply(message []byte) {
	for conn := range server.client {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (server *Server) SetUsername(connection *websocket.Conn, message []byte) string {
	username := strings.TrimLeft(string(message), "Username: ")
	username, found := strings.CutSuffix(username, " ")
	if !found && len(username) == 0 {
		log.Println(username)
		return "Username is blank, please try again"
	}
	server.client[connection] = username

	return "Username set as: " + username
}

func (server *Server) CheckForUsername(connection *websocket.Conn) bool {
	check := server.client[connection]
	return check != ""
}

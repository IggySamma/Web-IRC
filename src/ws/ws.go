package ws

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	sync.RWMutex
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
		sync.RWMutex{},
		make(map[*websocket.Conn]string),
		handleMessage,
	}

	channels := &Channel{
		sync.RWMutex{},
		make(map[string]*LinkedList),
	}

	err := channels.CreateChannel(string("Global"), string(""))
	if err != string("Channel added") {
		fmt.Println(err)
	}

	http.HandleFunc("/ws", server.WebsocketHandler)
	http.HandleFunc("/channels", server.Reply([]byte(channels.GetChannels())))

	go func() {
		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return &server
}

func MessageHandler(server *Server, connection *websocket.Conn, message []byte) {
	/*if(string(message) == "Request channels"){
		server.Reply()
	}*/
	if strings.Contains(string(message), "Username: ") {
		previous := server.client[connection]
		SetupUser(server, connection, message)
		if !(len(previous) == 0) {
			server.Reply([]byte(previous + string(" updated username to ") + server.client[connection]))
		}
	} else if server.CheckForClient(connection) && string(message) != "" {
		server.Reply([]byte(server.client[connection] + string(": ") + string(message)))
	}
}

func SetupUser(server *Server, connection *websocket.Conn, message []byte) {
	/*if !server.CheckForClient(connection) {
		connection.WriteMessage(websocket.TextMessage, []byte("Enter username"))
	} */
	if strings.Contains(string(message), "Username: ") {
		connection.WriteMessage(websocket.TextMessage, []byte(server.SetUsername(connection, message)))
	}
}

func (server *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := WebsocketBuffer.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	server.Lock()
	server.client[connection] = ""
	server.Unlock()

	messageType, username, err := connection.ReadMessage()
	if err != nil || messageType == websocket.CloseMessage {
		server.ClearUser(connection)
	}

	go SetupUser(server, connection, username)

	for {
		messageType, message, err := connection.ReadMessage()

		if err != nil || messageType == websocket.CloseMessage {
			fmt.Println(err)
			break
		}

		go server.handleMessage(server, connection, message)
	}

	server.ClearUser(connection)
}

func (server *Server) ClearUser(connection *websocket.Conn) {
	server.Lock()
	delete(server.client, connection)
	server.Unlock()
	connection.Close()
}

func (server *Server) Reply(message []byte) {
	fmt.Println(string(message))
	for conn := range server.client {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (server *Server) SetUsername(connection *websocket.Conn, message []byte) string {
	username := strings.TrimPrefix(string(message), "Username: ")
	username, found := strings.CutSuffix(username, " ")
	log.Println(string(message))
	if !found && len(username) == 0 {
		log.Println(username)
		return "Error: Username is blank/starts with a space. Please try again."
	} else if server.CheckForUsername(username) {
		fmt.Println(username + " already set")
		return "Error: Username already in use, please try another"
	} else {
		server.client[connection] = username
		fmt.Println(string("Added: ") + server.client[connection])
		return "Username set as: " + username
	}
}

func (server *Server) CheckForClient(connection *websocket.Conn) bool {
	server.RLock()
	check := server.client[connection]
	server.RUnlock()
	return check != ""
}

func (server *Server) CheckForUsername(username string) bool {
	server.RLock()
	for _, value := range server.client {
		if value == username {
			server.RUnlock()
			return true
		}
	}
	server.RUnlock()
	return false
}

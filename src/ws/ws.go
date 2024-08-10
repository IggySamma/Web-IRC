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
	channel       *Channel
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

	channels := &Channel{
		channel:  make(map[string]*LinkedList),
		password: make(map[string]uint32),
	}

	server := &Server{
		client:        make(map[*websocket.Conn]string),
		handleMessage: handleMessage,
		channel:       channels,
	}
	/* testing*/
	server.channel.AddChannel("Global 1", "", "test")
	server.channel.AddChannel("Global 2", "", "")
	server.channel.AddChannel("Global 3", "", "")
	server.channel.AddChannel("Global 4", "", "")
	server.channel.AddChannel("Global 5", "", "")

	http.HandleFunc("/ws", server.WebsocketHandler)

	go func() {
		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return server
}

func MessageHandler(server *Server, connection *websocket.Conn, message []byte) {
	fmt.Println("Raw Message: " + string(message))
	if string(message) == "/Request channels" {
		channels := server.channel.GetChannels()
		server.Reply(connection, channels)
	} else if strings.HasPrefix(string(message), "/Username:") {
		//else if strings.Contains(string(message), "Username: ") {
		previous := server.client[connection]
		go server.SetupUser(connection, string(message))
		if !(len(previous) == 0) {
			server.Reply(connection, previous+string(" updated username to ")+server.client[connection])
		}
	} else if strings.HasPrefix(string(message), "/Join:") {
		user := server.GetUsername(connection)
		reply := ""
		if user != "" {
			reply = server.channel.InserUserToChannel(strings.TrimPrefix(string(message), "/Join: "), user, "")
			server.Reply(connection, reply)
		}
		//server.Reply(connection, "Join channel: "+string(message))
	} else if strings.HasPrefix(string(message), "/Password:/Channel:") {
		user := server.GetUsername(connection)

		reply := server.channel.InserUserToChannel(
			MessageDelim(
				strings.TrimLeft(
					strings.TrimPrefix(
						string(message), "/Password:/Channel:"),
					":"),
				":",
				"Left"),
			user,
			MessageDelim(
				strings.TrimLeft(
					strings.TrimPrefix(
						string(message),
						"/Password:/Channel:"),
					":"),
				":",
				"Right"))

		if user != "" {
			server.Reply(connection, reply)
		}
	} else if server.CheckForClient(connection) && strings.HasPrefix(string(message), "/Channel:") {
		server.ReplyAll(MessageDelim(strings.TrimPrefix(string(message), "/Channel:"), ":", "Right"),
			server.client[connection]+": "+MessageDelim(strings.TrimPrefix(string(message), "/Channel:"), ":", "Left"))
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

func (server *Server) Reply(connection *websocket.Conn, message string) {
	fmt.Println(message)
	connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func (server *Server) ReplyAll(channel string, message string) {
	fmt.Println("Sending message to channel:", channel)
	list, exists := server.channel.channel[channel]

	fmt.Println(server.channel.channel)
	fmt.Println(server.channel.channel["Global 3"])

	if !exists {
		fmt.Println("Channel does not exist")
		return
	}

	if list == nil {
		fmt.Println("List is nil")
		return
	}

	list.RLock()
	defer list.RUnlock()

	users := GetUsersInChannel(list)
	if users == "" {
		fmt.Println("No users in the channel")
		return
	}

	fmt.Println("Users in channel:", users)

	for _, username := range strings.Split(users, ",") {
		username = strings.TrimSpace(username)
		if username != "" {
			conn := server.RetriveConnectionFromUsername(username)
			if conn != nil {
				if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Printf("Failed to send message to %s: %v", username, err)
				}
			}
		}
	}
}

/* Setting up user */

func (server *Server) CheckForClient(connection *websocket.Conn) bool {
	server.RLock()
	check := server.client[connection]
	server.RUnlock()
	return check != ""
}

func (server *Server) RetriveConnectionFromUsername(username string) (key *websocket.Conn) {
	for k, value := range server.client {
		if value == username {
			key = k
			return
		}
	}
	return
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

func (server *Server) GetUsername(connection *websocket.Conn) string {
	server.RLock()
	defer server.RUnlock()
	temp := server.client
	user, err := temp[connection]
	if !err {
		return ""
	}
	return user
}

func (server *Server) SetupUser(connection *websocket.Conn, message string) {
	/*if !server.CheckForClient(connection) {
		connection.WriteMessage(websocket.TextMessage, []byte("Enter username"))
	}*/
	if strings.HasPrefix(string(message), "/Username: ") {
		server.Reply(connection, server.SetUsername(connection, message))
	}
}

func (server *Server) ClearUser(connection *websocket.Conn) {
	server.Lock()
	delete(server.client, connection)
	server.Unlock()
	connection.Close()
}

func (server *Server) SetUsername(connection *websocket.Conn, message string) string {
	username := strings.TrimPrefix(message, "/Username: ")
	username, found := strings.CutSuffix(username, " ")
	log.Println(message)
	if !found && len(username) == 0 {
		log.Println(username)
		return "Username Error: Username is blank/starts with a space. Please try again."
	} else if server.CheckForUsername(username) {
		fmt.Println(username + " already set")
		return "Username Error: Username already in use, please try another"
	} else {
		server.client[connection] = username
		fmt.Println(string("Added: ") + server.client[connection])
		return "Username set as: " + username
	}
}

func MessageDelim(message string, delim string, direction string) string {
	if idx := strings.Index(message, delim); idx != -1 {
		if direction == "Right" {
			return message[:idx]
		} else {
			return message[idx+1:]
		}
	}
	return message
}

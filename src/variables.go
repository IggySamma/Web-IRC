package main

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var WebsocketBuffer = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//Temp for test remove later
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Usernames = make(map[net.Addr]string)

var Channels = make(map[int]string)

type Users struct {
	Usernames map[net.Addr]string
	/*access map[*websocket.Conn]int*/
}

package main

import (
	"github.com/IggySamma/Web-IRC/ws"
)

func main() {
	ws.StartServer(ws.MessageHandler)
	/* testing list
	list := &ws.LinkedList{}
	list.Insert("test1", "admin")
	list.Insert("test2", "user")

	fmt.Println(ws.GetUsersInChannel(list))

	list.DeleteUserFromChannel("test1")

	fmt.Println(ws.GetUsersInChannel(list))
	*/
	for {
	}
}

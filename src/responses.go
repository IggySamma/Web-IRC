package main

import (
	"log"
	"net"

	/*"strconv"*/

	/*"os"*/
	"strings"
)

func MessageHandler(connection net.Addr, message string) string {
	if message != "" {
		log.Println("Failed: " + message)
	}
	if !CheckForUsername(connection) {
		return "Enter username: "
	}
	if strings.Contains(message, "Username: ") {
		return SetUsername(connection, message)
	}
	return message
}

func SetUsername(connection net.Addr, message string) string {
	username := Usernames[connection]
	username = strings.TrimLeft(message, "Username: ")
	username, found := strings.CutSuffix(username, " ")
	if !found && len(username) == 0 {
		log.Println(username)
		return "Username is blank, please try again"
	}
	Usernames[connection] = username

	return "Username set as: " + username
}

func CheckForUsername(connection net.Addr) bool {
	_, check := Usernames[connection]
	return check
}

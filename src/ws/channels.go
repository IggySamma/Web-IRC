package ws

import (
	"fmt"
	"hash/fnv"
	"log"
	"strings"
	"sync"
)

type Node struct {
	user      string
	privilege string
	next      *Node
	previous  *Node
}

type LinkedList struct {
	sync.RWMutex
	head *Node
}

type Channel struct {
	sync.RWMutex
	channel  map[string]*LinkedList
	password map[string]uint32
}

func (list *LinkedList) Insert(user string, privileges string) {
	newNode := &Node{
		user:      user,
		privilege: privileges,
		next:      nil,
		previous:  nil,
	}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	newNode.previous = current
	current.next = newNode
}

func (list *LinkedList) DeleteUserFromChannel(user string) {
	var current *Node = list.head

	for current != nil && current.user != user {
		current = current.next
	}

	if current == nil {
		return
	}

	if current == list.head {
		list.head = current.next
		if list.head != nil {
			list.head.previous = nil
		}
		return
	}

	if current.next == nil {
		current.previous.next = nil
		return
	}

	current.previous.next = current.next
	current.next.previous = current.previous
}

func GetUsersInChannel(list *LinkedList) string {
	if list.head == nil {
		return ""
	}
	var result string
	current := list.head
	for current != nil {
		if result != "" {
			result += ", "
		}
		result += current.user
		current = current.next
	}
	return result
}

func (c *Channel) GetChannels() string {
	/*
		for channels := range c.channel {
			if channels != "" {
				return channels
			}
		}*/
	var channels []string
	for temp := range c.channel {
		channels = append(channels, temp)
	}

	return "Channels: " + strings.Join(channels, ",")
}

func (c *Channel) AddChannel(channelName string, username string, password string) string {
	if _, check := c.channel[channelName]; check {
		return "Channel already exists"
	}

	c.channel[channelName] = &LinkedList{head: nil}
	c.password[channelName] = HashPass(password)
	if username != "" {
		c.channel[channelName].head = &Node{
			user:      username,
			privilege: "Admin",
		}
	} else {
		c.channel[channelName].head = &Node{}
	}
	return "Channel added"
}

func (c *Channel) DeleteChannel(channelName string, privilege string) string {
	if privilege != string("Admin") {
		return string("Not enough privileges to delete the channel")
	}
	delete(c.channel, channelName)
	return string("Channel deleted")
}

func HashPass(input string) uint32 {
	temp := fnv.New32a()
	temp.Write([]byte(input))
	return temp.Sum32()
}

func (c *Channel) InserUserToChannel(channel string, user string, password string) string {
	c.Lock()
	defer c.Unlock()
	/*This is an IRC not a banking app so should be enough*/
	if c.password[channel] != HashPass("") && c.password[channel] != HashPass(password) {
		return "Enter password."
	}
	fmt.Println("1")
	exists := GetUsersInChannel(c.channel[channel])
	for _, username := range strings.Split(exists, ",") {
		username = strings.TrimSpace(username)
		if username == user {
			c.RemoveUserFromChannel(channel, user)
			log.Println(username + " removed from channel")
		}
	}

	list := c.channel[channel].head
	fmt.Println("2")
	if list == nil {
		return "Channel doesn't exist."
	}
	fmt.Println("3")
	for list.next != nil {
		list = list.next
	}
	fmt.Println("4")
	newNode := &Node{
		user:      user,
		privilege: "User",
		previous:  list,
	}
	list.next = newNode
	fmt.Println("5")

	return "Joined " + channel
}

func (c *Channel) RemoveUserFromChannel(channel string, user string) {
	list := c.channel[channel].head

	for list.user != user {
		if list.next != nil {
			list = list.next
		}
	}

	if list.user != user {
		return
	}

	if list == c.channel[channel].head {
		if list.next != nil {
			c.channel[channel].head = list.next
		}
		return
	}
	if list.next == nil {
		list.previous.next = nil
		return
	} else {
		list.previous.next = list.next
		list.next.previous = list.previous
	}

}

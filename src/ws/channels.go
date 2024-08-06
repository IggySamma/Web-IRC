package ws

import "sync"

/*
	type Channel struct {
		channel map[string]*LinkedList
	}
*/
type Node struct {
	user      string
	privilege string
	next      *Node
	previous  *Node
}

type LinkedList struct {
	head *Node
}

type Channel struct {
	sync.RWMutex
	channel map[string]*LinkedList
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
		if list.head.previous != nil {
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
	c.RLock()
	defer c.RUnlock()
	for channels, _ := range c.channel {
		if channels != "" {
			return channels
		}
	}
	return ""
}

func (c *Channel) CreateChannel(channelName string, username string) string {
	c.Lock()
	defer c.Unlock()
	if _, check := c.channel[channelName]; check {
		return string("Channel already exists")
	}

	c.channel[channelName] = &LinkedList{}
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

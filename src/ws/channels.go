package ws

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

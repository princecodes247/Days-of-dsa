package main


type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

// Creating a new node and inserting it at the beginning of the linked list.
func (c *CircularLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if c.head == nil {
		newNode.next = newNode
		c.head = newNode
		return
	}
	current := c.head
	for current.next != c.head {
		current = current.next
	}
	newNode.next = c.head
	current.next = newNode
	c.head = newNode
}
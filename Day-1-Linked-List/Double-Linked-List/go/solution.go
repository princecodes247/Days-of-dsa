package main

import "fmt"


type Node struct {
	data int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
}

// Creating a new node and inserting it at the beginning of the linked list.
func (d *DoubleLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = d.head
	newNode.prev = nil
	if d.head != nil {
		d.head.prev = newNode
	}
	d.head = newNode
}

// Creating a new node and inserting it at the end of the linked list.
func (d *DoubleLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if d.head == nil {
		newNode.prev = nil
		d.head = newNode
		return
	}
	last := d.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
	newNode.prev = last
}

// Creating a new node and inserting it at the given position of the linked list.
func (d *DoubleLinkedList) InsertAtPosition(data, position int) {
	newNode := &Node{data: data}
	if position == 0 {
		newNode.next = d.head
		newNode.prev = nil
		if d.head != nil {
			d.head.prev = newNode
		}
		d.head = newNode
		return
	}
	current := d.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil {
		fmt.Println("Position is out of range")
		return
	}
	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
}

// Deleting the first node of the linked list.
func (d *DoubleLinkedList) DeleteFirstNode() {
	if d.head == nil {
		fmt.Println("List is empty")
		return
	}
	d.head = d.head.next
	d.head.prev = nil
}

// Deleting the last node of the linked list.
func (d *DoubleLinkedList) DeleteLastNode() {
	if d.head == nil {
		fmt.Println("List is empty")
		return
	}
	last := d.head
	for last.next != nil {
		last = last.next
	}
	last.prev.next = nil
}

// Deleting the node at the given position of the linked list.
func (d *DoubleLinkedList) DeleteNodeAtPosition(position int) {
	if d.head == nil {
		fmt.Println("List is empty")
		return
	}
	if position == 0 {
		d.head = d.head.next
		d.head.prev = nil
		return
	}
	current := d.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil || current.next == nil {
		fmt.Println("Position is out of range")
		return
	}
	current.next = current.next.next
	current.next.prev = current
}

// Printing the linked list.
func (d *DoubleLinkedList) PrintList() {
	if d.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := d.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	dll := &DoubleLinkedList{}
	dll.InsertAtEnd(1)
	dll.InsertAtEnd(2)
	dll.InsertAtEnd(3)
	dll.InsertAtEnd(4)
	dll.InsertAtEnd(5)
	dll.PrintList()
	dll.InsertAtBeginning(0)
	dll.PrintList()
	dll.InsertAtPosition(6, 3)
	dll.PrintList()
	dll.DeleteFirstNode()
	dll.PrintList()
	dll.DeleteLastNode()
	dll.PrintList()
	dll.DeleteNodeAtPosition(2)
	dll.PrintList()
}
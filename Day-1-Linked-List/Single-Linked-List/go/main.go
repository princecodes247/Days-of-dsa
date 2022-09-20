package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type SingleLinkedList struct {
	head *Node
}

// Creating a new node and inserting it at the beginning of the linked list.
func (s *SingleLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = s.head
	s.head = newNode
}

// Creating a new node and inserting it at the end of the linked list.
func (s *SingleLinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if s.head == nil {
		s.head = newNode
		return
	}
	last := s.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

// Creating a new node and inserting it at the given position of the linked list.
func (s *SingleLinkedList) InsertAtPosition(data, position int) {
	newNode := &Node{data: data}
	if position == 0 {
		newNode.next = s.head
		s.head = newNode
		return
	}
	current := s.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil {
		fmt.Println("Position is out of range")
		return
	}
	newNode.next = current.next
	current.next = newNode
}

// Deleting the first node of the linked list.
func (s *SingleLinkedList) DeleteFirstNode() {
	if s.head == nil {
		fmt.Println("List is empty")
		return
	}
	s.head = s.head.next
}

// Deleting the last node of the linked list.
func (s *SingleLinkedList) DeleteLastNode() {
	if s.head == nil {
		fmt.Println("List is empty")
		return
	}
	if s.head.next == nil {
		s.head = nil
		return
	}
	last := s.head
	secondLast := s.head
	for last.next != nil {
		secondLast = last
		last = last.next
	}
	secondLast.next = nil
}

// Deleting the node at the given position of the linked list.
func (s *SingleLinkedList) DeleteNodeAtPosition(position int) {
	if s.head == nil {
		fmt.Println("List is empty")
		return
	}
	if position == 0 {
		s.head = s.head.next
		return
	}
	current := s.head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil || current.next == nil {
		fmt.Println("Position is out of range")
		return
	}
	current.next = current.next.next
}

// Printing the linked list.
func (s *SingleLinkedList) PrintList() {
	current := s.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	singleLinkedList := SingleLinkedList{}
	singleLinkedList.InsertAtBeginning(1)
	singleLinkedList.InsertAtBeginning(2)
	singleLinkedList.InsertAtBeginning(3)
	singleLinkedList.InsertAtBeginning(4)
	singleLinkedList.InsertAtBeginning(5)
	singleLinkedList.PrintList()
	singleLinkedList.InsertAtEnd(6)
	singleLinkedList.PrintList()
	singleLinkedList.InsertAtPosition(7, 3)
	singleLinkedList.PrintList()
	singleLinkedList.DeleteFirstNode()
	singleLinkedList.PrintList()
	singleLinkedList.DeleteLastNode()
	singleLinkedList.PrintList()
	singleLinkedList.DeleteNodeAtPosition(2)
	singleLinkedList.PrintList()
}
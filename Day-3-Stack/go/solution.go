package main

import "fmt"

type Stack struct {
	size int
	data []int
}

func (s *Stack) Push(data int) int {
	if len(s.data) == s.size {
		fmt.Println("Stack is full")
		return -1 // Stack is full
	}

	s.data = append(s.data, data)
	return data
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return data
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	stack := Stack{3, []int{}}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
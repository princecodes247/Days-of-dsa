package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) Rear() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[len(q.data)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	fmt.Println(q.Size())
	fmt.Println(q.IsEmpty())
}
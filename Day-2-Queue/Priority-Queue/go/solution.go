package main

import "fmt"

type QueueItem struct {
	priority int
	value    int
}

type PriorityQueue struct {
	data []QueueItem
}

func (q *PriorityQueue) Enqueue(priority, value int) {
	q.data = append(q.data, QueueItem{priority, value})

	for i := len(q.data) - 1; i > 0; i-- {
		if q.data[i].priority < q.data[i-1].priority {
			q.data[i], q.data[i-1] = q.data[i-1], q.data[i]
		}
	}
}

func (q *PriorityQueue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0].value
	q.data = q.data[1:]
	return val
}

func (q *PriorityQueue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0].value
}

func (q *PriorityQueue) Rear() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[len(q.data)-1].value
}

func (q *PriorityQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *PriorityQueue) Size() int {
	return len(q.data)
}

func main() {
	q := PriorityQueue{}
	q.Enqueue(1, 1)
	q.Enqueue(2, 2)
	q.Enqueue(3, 3)
	q.Enqueue(4, 4)
	q.Enqueue(5, 5)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	fmt.Println(q.Size())
	fmt.Println(q.IsEmpty())
}
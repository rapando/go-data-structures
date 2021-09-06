package main

import (
	"fmt"
	"github.com/rapando/go-data-structures/queues"
)

func main() {
	var q *queues.Queue
	var err error

	q, err = queues.InitializeQueue(4)
	if err != nil {
		fmt.Println(err)
		return
	}

	q.Enqueue(2)
	f, _ := q.Peek()
	fmt.Println("Queued :", 2, "Peek : ", f)

	q.Enqueue(4)
	f, _ = q.Peek()
	fmt.Println("Queued :", 4, "Peek : ", f)

	q.Enqueue(6)
	f, _ = q.Peek()
	fmt.Println("Queued :", 6, "Peek : ", f)


	i, _ := q.Dequeue()
	fmt.Println("Dequeued : ", i)

	i, _ = q.Dequeue()
	fmt.Println("Dequeued : ", i)

	i, _ = q.Dequeue()
	fmt.Println("Dequeued : ", i)

	i, _ = q.Dequeue()
	fmt.Println("Dequeued : ", i)
}
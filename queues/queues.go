package queues

import "fmt"

type Queue struct {
	size int
	front interface{}
	items []interface{}
}

// Initialize : creates a new queue instance
func InitializeQueue(size int) (*Queue, error) {
	if size <= 1 {
		return nil, fmt.Errorf("queue size must be greater than 1")
	}
	return &Queue {
		size: size,
		front: nil,
		items: nil,
	}, nil
}

// Size : returns the size of the current queue
func (q *Queue) Size() int {
	return len(q.items)
}

// isEmpty : returns true if queue is empty, false if otherwise
func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

// isFull : returns true if queue is full, false if otherwise
func (q *Queue) isFull() bool {
	return int(q.size) == q.Size()
}

// Enqueue : Add an item to the front of the queue
func (q *Queue) Enqueue(item interface{}) (err error) {
	if q.isFull() {
		return fmt.Errorf("queue is full")
	}
	q.items = append([]interface{}{item}, q.items...)
	q.front = q.items[q.Size() -1]
	return nil
}

// Dequeue : Remove the item at the back of the queue
func (q *Queue) Dequeue() (item interface{}, err error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	currentSize := q.Size()
	if currentSize == 1 {
		item = q.items[0]
		q.front = nil
		q.items = nil
		return item, nil
	}

	item = q.items[currentSize - 1]
	q.items = q.items[0 : currentSize - 1]
	q.front, err = q.Peek()
	return item, err
}

// Peek : get the item at the front of the queue without removing it
func (q *Queue) Peek() (item interface{}, err error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.front, nil
}
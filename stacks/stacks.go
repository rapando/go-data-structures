package stacks

import (
	"fmt"
)

type Stack struct {
	size  int
	top   interface{}
	items []interface{}
}

// Initialize : creates a new stack instance
func InitializeStack(size int) (*Stack, error) {
	if size <= 1 {
		return nil, fmt.Errorf("stack size must be greater than 1")
	}
	return &Stack{
		size: size,
		top:  nil,
	}, nil
}

// Size : returns the size of the current stack
func (s *Stack) Size() int {
	return len(s.items)
}

// isEmpty : returns true if stack is empty, false if otherwise
func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

// isFull : returns true if stack is full, false if otherwise
func (s *Stack) isFull() bool {
	return int(s.size) == s.Size()
}


// Push : adds an item to the top of the stack stack
func (s *Stack) Push(item interface{}) (err error) {
	if s.isFull() {
		return fmt.Errorf("stack is full")
	}
	s.top = item
	s.items = append(s.items, item)
	return nil
}

// Pop : removes an item from the top of the stack
func (s *Stack) Pop() (item interface{}, err error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	currentSize := s.Size()

	if currentSize == 1 {
		item = s.items[0]
		s.top = nil
		s.items = nil
		return item, nil
	}

	item = s.items[currentSize-1]
	s.items = s.items[0 : currentSize-1]
	s.top, err = s.Peek()
	return item, err
}

// Peek : returns the top item
func (s *Stack) Peek() (item interface{}, err error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.top, nil
}

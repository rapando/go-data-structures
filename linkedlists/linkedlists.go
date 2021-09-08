package linkedlists

import "fmt"

type LinkedList struct {
	data interface{}
	next *LinkedList
}

// Initialize : creates a new linked list
func InitializeLinkedList(data interface{}) *LinkedList {
	return &LinkedList{
		data: data,
		next: nil,
	}
}

// GetTail : get tail node()
func (node *LinkedList) GetTail() *LinkedList {
	for {
		if node.next != nil {
			node = node.next
		} else {
			return node
		}
	}
}

// AddNode : adds a node to another node
func (node *LinkedList) AddNode(data interface{}) {
	tail := node.GetTail()
	tail.next = &LinkedList{
		data: data,
	}
}

// Traverse: traverse through the linked list
func (node *LinkedList) Traverse() (list string) {
	for node.next != nil {
		list += fmt.Sprintf("%v ", node.data)
		node = node.next
	}
	list += fmt.Sprintf("%v", node.data)
	return list
}

// Search : search for data in a linked list
func (node *LinkedList) Search(item interface{}) (found bool, position int) {
	for node.next != nil {
		if node.data == item {
			return true, position
		}
		node = node.next
		position++
	}
	if node.data == item {
		return true, position
	}
	return false, 0
}

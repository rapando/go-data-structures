package linkedlists

import "fmt"

type LinkedList struct {
	length int
	data   int
	next   *LinkedList
}

// Initialize : creates a new linked list
func InitializeLinkedList(data int) *LinkedList {
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
func (node *LinkedList) AddNode(data int) {
	tail := node.GetTail()
	tail.next = &LinkedList{
		data: data,
	}
}

// Traverse: traverse through the linked list
func (node *LinkedList) Traverse() (list string) {
	for node.next != nil {
		list += fmt.Sprintf("%d ", node.data)
		node = node.next
	}
	list += fmt.Sprintf("%d", node.data)
	return list
}

// Search : search for data in a linked list
func (node *LinkedList) Search(item int) (found bool, position int) {
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

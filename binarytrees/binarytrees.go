package binarytrees

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  int
}

type BinaryTree struct {
	Root *BinaryNode
}

// InitializeBinaryTree - initializes a new binary tree
// takes in data to be assigned to the root
// it inserts the smaller value to the Left and the greater value to Right
// returns the binary tree
func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

func (n *BinaryNode) Insert(data int) {
	if n == nil {
		return
	}

	if data <= n.Data {
		if n.Left == nil {
			n.Left = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(data)
		}
	}
}

func PrintBinaryTree(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	PrintBinaryTree(w, node.Left, ns+2, 'L')
	PrintBinaryTree(w, node.Right, ns+2, 'R')
}

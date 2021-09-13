package main

import (
	"os"

	"github.com/rapando/go-data-structures/binarytrees"
)

func main() {
	tree := &binarytrees.BinaryTree{}
	tree.
		Insert(10).
		Insert(9).
		Insert(11).
		Insert(8).
		Insert(7)
	binarytrees.PrintBinaryTree(os.Stdout, tree.Root, 0, 'M')
}

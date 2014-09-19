package tree

import (
	"fmt"
	"testing"
)

func Test_binary_tree(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Add(3)
	tree.Add(2)
	tree.Add(5)
	tree.Add(1)
	tree.Add(4)
	tree.Add(6)
	tree.Traverse()
	fmt.Println()
	tree.Search(5).TraverseByLevel()
	fmt.Println()
	tree.Remove(3)
	tree.TraverseByLevel()
	fmt.Println()
}
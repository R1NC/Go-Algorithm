package main

import "fmt"

type BinaryTree struct {
	Value int
	LeftChild *BinaryTree
	RightChild *BinaryTree
}

func (tree *BinaryTree) Add(value int) {
	if tree.Value == 0 {
		tree.Value = value
		tree.LeftChild = &BinaryTree{}
		tree.RightChild = &BinaryTree{}
        } else {
		if value < tree.Value {
			tree.LeftChild.Add(value)
		} else {
			tree.RightChild.Add(value)
		}
	}
}

func (tree *BinaryTree) Search(value int) *BinaryTree {
	if tree == nil {
		return nil
	}
	if tree.Value == value {
		return tree
	} else {
		if value < tree.Value {
			return tree.LeftChild.Search(value)
		} else {
			return tree.RightChild.Search(value)
		}
	}
}

func (tree *BinaryTree) Print() {
	fmt.Printf("%d", tree.Value)
	if tree.LeftChild == nil && tree.RightChild == nil {
		return
	}
	fmt.Print("(")
	tree.LeftChild.Print()
	fmt.Print(",")
	tree.RightChild.Print()
	fmt.Print(")")
}

func main() {
	tree := &BinaryTree{}
	tree.Add(3)
	tree.Add(2)
	tree.Add(5)
	tree.Add(1)
	tree.Add(4)
	tree.Add(6)
	tree.Print()
	fmt.Println()
	tree.Search(5).Print()
	fmt.Println()
}

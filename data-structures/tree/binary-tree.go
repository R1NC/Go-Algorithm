package main

import "fmt"

type BinaryTree struct {
	value interface{}
	left_child *BinaryTree
	right_child *BinaryTree
}

func (tree *BinaryTree) Add(value interface{}) {
	if tree == nil {
        tree.value = value
        tree.left_child  = nil
        tree.right_child = nil
    } else {
        if value < tree.value {
            tree.left_child.Insert(value)
        } else {
            tree.right_child.Insert(value)
        }
    }
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.value == nil {
        return nil
    }
    if tree.value == value {
        return tree
    } else {
        if value < tree.value {
            return tree.left_child.Search(value)
        } else {
            return tree.right_child.Search(value)
        }
    }
}

func main() {
	tree := &BinaryTree{}
	tree.Add(3)
	tree.Add(2)
	tree.Add(5)
	tree.Add(1)
	tree.Add(4)
	tree.Add(6)
	t := tree.Search(5)
}
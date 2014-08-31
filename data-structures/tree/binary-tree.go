package tree

import (
    "fmt"
    "Go-Algorithm/data-structures/queue"
)

type BinaryTree struct {
	Value int
	LeftChild *BinaryTree
	RightChild *BinaryTree
	size int
}

func (tree *BinaryTree) Size() int {
	return tree.size
}

func (tree *BinaryTree) Add(value int) {
	if tree.Value == 0 {
		tree.Value = value
		tree.LeftChild = &BinaryTree{}
		tree.RightChild = &BinaryTree{}
		tree.size++
	} else {
		if value < tree.Value {
			tree.LeftChild.Add(value)
		} else {
			tree.RightChild.Add(value)
		}
	}
}

func (tree *BinaryTree) Remove(value int) {
	t := tree.Search(value)
	if t == nil {
		panic("This value is not found.")
	}
	if t.LeftChild.Value == 0 && t.RightChild.Value == 0 {// Both LeftChild and RightChild are empty.
		// Remove it directly.
		destroy(t)
	} else if t.LeftChild.Value != 0 && t.RightChild.Value != 0 {// Both LeftChild and RightChild are not empty.
		// Get the min-element in its RightChild.
		min := t.RightChild.getMin()
		// Replace its value with the min-element's value.
		t.Value = min.Value
		// Remove the min-element.
		destroy(min)
	} else {// LeftChild or RightChild is empty.
		// Replace it with its child which is not empty.
		if t.LeftChild.Value > 0 {
			t.replaceWith(t.LeftChild)
		} else {
			t.replaceWith(t.RightChild)
		}
	}
	tree.size--
}

func destroy(tree *BinaryTree) {
	tree.Value = 0
	tree.LeftChild = nil
	tree.RightChild = nil
}

func (tree *BinaryTree) getMin() *BinaryTree {
	if tree.LeftChild.Value == 0 {
		return tree
	} else {
		return tree.LeftChild.getMin()
	}
}

func (tree *BinaryTree) replaceWith(t *BinaryTree) {
	tree.Value = t.Value
	tree.LeftChild = t.LeftChild
	tree.RightChild = t.RightChild
	destroy(t)
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

func (tree *BinaryTree) Traverse() {
	fmt.Printf("%d", tree.Value)
	if tree.LeftChild == nil && tree.RightChild == nil {
		return
	}
	fmt.Print("(")
	tree.LeftChild.Traverse()
	fmt.Print(",")
	tree.RightChild.Traverse()
	fmt.Print(")")
}

func convertInterfaceToBinaryTreePointer(x interface{}) *BinaryTree {
    if v, ok := x.(*BinaryTree); ok {
        return v
    } else {
        panic("Type convertion exception.")
    }
}

func (tree *BinaryTree) TraverseByLevel() {
    queue := &queue.LinkedQueue{}
    queue.Add(tree)
    for queue.Size() > 0 {
        t := convertInterfaceToBinaryTreePointer(queue.Peek())
        if t.LeftChild != nil {
            queue.Add(t.LeftChild)
        }
        if t.RightChild != nil {
            queue.Add(t.RightChild)
        }
        fmt.Printf("%d ", t.Value)
        queue.Remove()
    }
}

/*
func main() {
	tree := &BinaryTree{}
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
	fmt.Println("Size: ", tree.Size())
	tree.Remove(3)
	tree.TraverseByLevel()
	fmt.Println()
	fmt.Println("Size: ", tree.Size())
}
*/
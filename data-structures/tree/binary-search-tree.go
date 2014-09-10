package tree

import (
    "fmt"
    "github.com/RincLiu/Go-Algorithm/data-structures/queue"
)

type BinarySearchTree struct {
	Value int
	LeftChild *BinarySearchTree
	RightChild *BinarySearchTree
}

func (tree *BinarySearchTree) Add(value int) {
	if tree.Value == 0 {
		tree.Value = value
		tree.LeftChild = &BinarySearchTree{}
		tree.RightChild = &BinarySearchTree{}
	} else {
		if value < tree.Value {
			tree.LeftChild.Add(value)
		} else {
			tree.RightChild.Add(value)
		}
	}
}

func (tree *BinarySearchTree) Remove(value int) {
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
}

func destroy(tree *BinarySearchTree) {
	tree.Value = 0
	tree.LeftChild = nil
	tree.RightChild = nil
}

func (tree *BinarySearchTree) getMin() *BinarySearchTree {
	if tree.LeftChild.Value == 0 {
		return tree
	} else {
		return tree.LeftChild.getMin()
	}
}

func (tree *BinarySearchTree) replaceWith(t *BinarySearchTree) {
	tree.Value = t.Value
	tree.LeftChild = t.LeftChild
	tree.RightChild = t.RightChild
	destroy(t)
}

func (tree *BinarySearchTree) Search(value int) *BinarySearchTree {
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

func (tree *BinarySearchTree) Traverse() {
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

func convertToBinarySearchTree(x interface{}) *BinarySearchTree {
    if v, ok := x.(*BinarySearchTree); ok {
        return v
    } else {
        panic("Type convertion exception.")
    }
}

func (tree *BinarySearchTree) TraverseByLevel() {
    queue := &queue.LinkedQueue{}
    queue.Add(tree)
    for queue.Size() > 0 {
        t := convertToBinarySearchTree(queue.Peek())
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
*/

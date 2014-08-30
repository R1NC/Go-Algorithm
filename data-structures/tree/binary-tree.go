package tree

import (
    "fmt"
    "Go-Algorithm/data-structures/queue"
)

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

func convertInterfaceToBinaryTreePointer(x interface{}) *BinaryTree {
    if v, ok := x.(*BinaryTree); ok {
        return v
    } else {
        panic("Type convertion exception.")
    }
}

func (tree *BinaryTree) PrintByLevel() {
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
<<<<<<< HEAD
    tree.Print()
    fmt.Println()
	tree.Search(5).PrintByLevel()
    fmt.Println()
}
*/
=======
	tree.Print()
	fmt.Println()
	tree.Search(5).Print()
	fmt.Println()
}
>>>>>>> 9b096b56a0dda3796bef903dae5b5a978e05471a

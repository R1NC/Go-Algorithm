package stack

import (
	"fmt"
	"testing"
)

func Test_linked_stack(t *testing.T) {
	stack := &LinkedStack{}
	stack.Push("iOS")
	stack.Push("Android")
	stack.Push("WindowsPhone")
	stack.Push("Symbian")
	stack.Push("BalckBerry")
	for stack.Size() > 0 {
		fmt.Printf("%s\n", stack.Peek())
		stack.Pop()
	}
}
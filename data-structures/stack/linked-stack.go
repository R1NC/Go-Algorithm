package main

import "fmt"

type node struct {
	value interface{}
	next *node
}

type LinkedStack struct {
	top *node
	size int
}

func (stack *LinkedStack) Size() int {
	return stack.size
}

func (stack *LinkedStack) Push(value interface{}) {
	new_node := &node{value, stack.top}
	stack.top = new_node
	new_node = nil
	stack.size++
}

func (stack *LinkedStack) Pop() {
	if stack.size == 0 {
		panic("Stack is empty.")
	}
	top_node := stack.top
	stack.top = top_node.next
	top_node.value = nil
	top_node.next = nil
	top_node = nil
	stack.size--
}

func (stack *LinkedStack) Peek() interface{} {
	if stack.size == 0 {
		panic("Stack is empty.")
	}
	return stack.top.value
}

func main() {
	stack := &LinkedStack{}
	stack.Push("iOS")
	stack.Push("Android")
	stack.Push("WindowsPhone")
	stack.Push("Symbian")
	stack.Push("BalckBerry")
	for stack.Size() > 0 {
		fmt.Printf("%s", stack.Peek())
		stack.Pop()
	}
}

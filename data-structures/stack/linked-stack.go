package stack

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
	stack.top = &node{value, stack.top}
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
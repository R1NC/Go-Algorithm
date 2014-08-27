package main

import "fmt"

type node struct {
	value interface{}
	prev *node
	next *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Reverse() {
	if list.tail == nil {
		panic("Empty list.")
	}
	current_node := list.tail
	list.head = current_node
	for current_node.prev != nil {
		current_node.prev = current_node.next
		current_node.next = current_node.prev
		current_node = current_node.prev
	}
	list.tail = current_node
}

func (list *LinkedList) IndexOf(value interface{}) int {
	if list.head == nil {
		return -1
	}
	current_node := list.head
	index := 0
	for current_node.next != nil {
		if value == current_node.value {
			return index
		} else {
			current_node = current_node.next
			index++
		}
	}
	return -1
}

func (list *LinkedList) getNodeAt(index int) *node{
	current_node := nil
	if index + 1 <= list.size / 2 {
		current_node = list.head
		for i := 0; i < index; i++ {
			current_node = current_node.next
		}
	} else {
		current_node = list.tail
		for i := list.size - 1; i > index; i-- {
			current_node = current_node.prev
		}
	}
	return current_node
}

func (list *LinkedList) Get(index int) interface{} {
	if index < 0 || (index >= list.size && list.size > 0) {
		panic("Invalid index.")
	}
	if list.head == nil {
		panic("Empty list.")
	}
	return list.getNodeAt(index).value
}

func (list *LinkedList) GetFirst() interface{} {
	if list.head == nil {
		panic("Empty list.")
	}
	return list.head.value
}

func (list *LinkedList) GetLast() interface{} {
	if list.tail == nil {
		panic("Empty list.")
	}
	return list.tail.value
}

func (list *LinkedList) Add(value interface{}, index int) {
	if index < 0 || (index > list.size && list.size > 0) {
		panic("Invalid index.")
	} else if index == 0 {
		list.AddToFirst(value)
	} else if index == list.size {
		list.AddToLast(value)
	} else {
		if list.head == nil {
			panic("Empty list and invalid index.")
		} else {
			new_node := &node{value, nil, nil}
			target_node := list.getNodeAt(index)
			prev_node := target_node.prev
			next_node := target_node
			prev_node.next = new_node
			new_node.prev = prev_node
			new_node.next = next_node
			next_node.prev = new_node
			list.size++
		}
	}
}

func (list *LinkedList) AddToFirst(value interface{}) {
	new_node := &node{value, nil, list.head}
	if list.head == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		list.head.prev = new_node
		list.head = new_node
	}
	list.size++
}

func (list *LinkedList) AddToLast(value interface{}) {
	new_node := &node{value, list.tail, nil}
	if list.tail == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		list.tail.next = new_node
		list.tail = new_node
	}
	list.size++
}

func (list *LinkedList) RemoveAt(index int) {
	if index < 0 || (index > list.size && list.size > 0) {
		panic("Invalid index.")
	}
	if index == 0 {
		list.RemoveFirst()
	} else if index == list.size {
		list.RemoveLast()
	} else {
		if list.size == 0 {
			panic("Empty list and invalid index.")
		} else {
			target_node := list.getNodeAt(index)
			prev_node := target_node.prev
			next_node := target_node.next
			target_node.prev = nil
			target_node.next = nil
			prev_node.next = next_node
			next_node.prev = prev_node
			list.size--
		}
	}
}

func (list *LinkedList) RemoveFirst() {
	if list.head == nil {
		panic("Empty list.")
	}
	first_node := list.head
	list.head = first_node.next
	first_node.next = nil
	list.size--
}

func (list *LinkedList) RemoveLast() {
	if list.tail == nil {
		panic("Empty list.")
	}
	last_node := list.tail
	list.tail = last_node.prev
	last_node.prev = nil
	list.size--
}

// Test:
func main() {
	list := &LinkedList{}
	list.AddToLast("D")
	list.AddToLast("E")
	list.AddToFirst("C")
	list.AddToFirst("B")
	list.AddToFirst("A")
	list.Add("G", 5)
	list.Add("F", 5)
	list.RemoveFirst()
	list.RemoveLast()
	list.RemoveAt(2)
	list.Reverse()
	size := list.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("%d : %s\n", i, list.Get(i))
	}
	fmt.Printf("Index of 'E': %d\n", list.IndexOf("E"))
	fmt.Printf("First: %s\n", list.GetFirst())
	fmt.Printf("Last: %s\n", list.GetLast())
}
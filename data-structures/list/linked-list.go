package main

import (
	"fmt"
)

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

func (list *LinkedList) Get(index int) interface{} {
	if index < 0 || (index >= list.size && list.size > 0) {
		panic("Invalid index.")
	}
	if list.head == nil {
		return nil
	}
	current_node := list.head
	for i := 0; i < index; i++ {
		current_node = current_node.next
	}
	return current_node.value
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

func (list *LinkedList) GetFirst() interface{} {
	if list.head == nil {
		return nil
	}
	return list.head.value
}

func (list *LinkedList) GetLast() interface{} {
	if list.tail == nil {
		return nil
	}
	return list.tail.value
}

// 'index' should start with 0;
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
			current_node := list.head
			for i := 0; i < index; i++ {
				current_node = current_node.next
			}
			prev_node := current_node.prev
			next_node := current_node
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
			current_node := list.head
			for i := 0; i < index; i++ {
        		current_node = current_node.next
			}
			prev_node := current_node.prev
			next_node := current_node.next
			current_node.prev = nil
			current_node.next = nil
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
	size := list.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("%d : %s\n", i, list.Get(i))
	}
	fmt.Printf("Index of 'E': %d\n", list.IndexOf("E"))
	fmt.Printf("First: %s\n", list.GetFirst())
	fmt.Printf("Last: %s\n", list.GetLast())
}
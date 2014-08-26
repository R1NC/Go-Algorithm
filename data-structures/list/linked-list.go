package list

type struct node {
	value interface{}
	prev *node
	next *node
}

type strcut LinkedList {
	head *node
	tail *node
	size int
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) GetValue(int index) interface{} {
	if list.head == nil {
		return nil
	}
	current_node := list.head
	for i := 0; i < index; i++ {
		current_node = current_node.next
	}
	return current_node.value
}

func (list *LinkedList) GetIndex(value interface{}) int {
	if list.head == nil {
		return -1
	}
	current_node := list.head
	index := 0
	for current_node.next != nil {
		if value == current_node.value {
			return index
		}
		else {
			current_node = current_node.next
			index++
		}
	}
	return -1
}

func (list *LinkedList) GetFirst() interface{} {
	return list.head.value
}

func (list *LinkedList) GetLast() interface{} {
	return list.tail.value
}

// 'index' should start with 0;
func (list *LinkedList) Add(value interface{}, index int) {
	if index < 0 || index >= list.size {
		return
	}
	new_node := &node{value, nil, nil}
	current_node := list.head
	for int i = 0; i <= index; i++ {
		current_node = current_node.next
	}
	prev_node := current_node
	next_node := current_node.next
	prev_node.next = new_node
	new_node.prev = prev_node
	new_node.next = next_node
	next_node.prev = new_node
	size++
}

func (list *LinkedList) Add(value interface{}) {
	new_node := &node{value, list.tail, nil}	
	list.tail.next = new_node
	list.tail = new_node
	size++
}

func (list *LinkedList) AddToFirst(value interface{}) {
	new_node := &node{value, nil, list.head}
	list.head.prev = new_node
	list.head = new_node
	size++
}

func (list *LinkedList) AddToLast(value interface{}) {
	list.Add(value)
}

func (list *LinkedList) Remove(index int) {
	if index < 0 || index >= list.size {
		return
	}
	current_node := list.head
	for int i = 0; i <= index; i++ {
        	current_node = current_node.next
	}
	prev_node := current_node.prev
	next_node := current_node.next
	current_node.prev = nil
	current_node.next = nil
	prev_node.next = next_node
	next_node.prev = prev_node
	size--	
}

func (list *LinkedList) RemoveFirst() {
	first_node := list.head
	list.head = first_node.next
	first_node.next = nil
	size--
}

func (list *LinkedList) RemoveLast() {
	last_node := list.tail
	list.tail = last_node.prev
	last_node.prev = nil
	size--
}

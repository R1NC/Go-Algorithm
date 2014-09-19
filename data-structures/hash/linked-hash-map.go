package hash

import "github.com/RincLiu/Go-Algorithm/data-structures/stack"

const TABLE_SIZE = 128

type entry struct {
	key int
	value interface{}
	next *entry
}

type LinkedHashMap struct {
	table []*entry
}

func (hashMap *LinkedHashMap) Put(key int, value interface{}) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e == nil {
		hashMap.table[hash] = &entry{key, value, nil}
	} else {
		for e.next != nil {
			if e.key == key {
				e.value = value
				return
			}
			e = e.next
		}
		if e.key == key {
			e.value = value
		} else {
			e.next = &entry{key, value, nil}
		}
	}
}

func (hashMap *LinkedHashMap) Get(key int) interface{} {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e != nil {
		for e != nil {
			if e.key == key {
				return e.value
			}
			e = e.next
		}
	}
	return nil
}

func (hashMap *LinkedHashMap) Remove(key int) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e != nil {
		for e.next != nil {
			if e.key == key {
				e.key = e.next.key
				e.value = e.next.value
				e.next = e.next.next
				return
			}
			e = e.next
		}
		if e.key == key {
			e = nil
		}
	}
}

func (hashMap *LinkedHashMap) Clear() {
	hashMap.checkTable()
	for _, e := range hashMap.table {
		if e != nil {
			stack := &stack.LinkedStack{}
			stack.Push(e)
			nextEntry := e.next
			for nextEntry != nil {
				stack.Push(nextEntry)
				nextEntry = nextEntry.next
			}
			for stack.Size() > 0 {
				e := convertToEntry(stack.Peek())
				e.key = 0
				e.value = nil
				e.next = nil
				e = nil
				stack.Pop()
			}
		}
	}
}

func getHashValue(key int) int {
	return key % TABLE_SIZE
}

func convertToEntry(x interface{}) *entry {
	if v, ok := x.(*entry); ok {
		return v
	} else {
		panic("Entry convertion exception.")
	}
}

func (hashMap *LinkedHashMap) checkTable() {
	if hashMap.table == nil {
		hashMap.table = []*entry{}
		for i := 0; i < TABLE_SIZE; i++ {
			hashMap.table = append(hashMap.table, nil)
		}
	}
}
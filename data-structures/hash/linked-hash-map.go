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
		for e != nil {
			if e.key == key {
				e.value = value
				return
			}
			e = e.next
		}
		e = &entry{key, value, nil}
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
		for e != nil {
			if e.key == key {
				nextEntry := e.next 
				e.value = nextEntry.value
				e.next = nextEntry.next
				nextEntry.key = 0
				nextEntry.value = nil
				nextEntry.next = nil
				return
			}
			e = e.next
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

/*
func main() {
	hashMap := &LinkedHashMap{}
	hashMap.Put(9, "android")
	hashMap.Put(9, "Android")
	hashMap.Put(137, "ios")
	hashMap.Put(137, "iOS")
	hashMap.Put(99, "windowsphone")
	hashMap.Put(99, "WindowsPhone")
	hashMap.Put(227, "symbian")
	hashMap.Put(227, "Symbian")
	hashMap.Put(355, "blackberry")
	hashMap.Put(355, "BlackBerry")
	fmt.Printf("9:%v\n", hashMap.Get(9))
	fmt.Printf("137:%v\n", hashMap.Get(137))
	fmt.Printf("99:%v\n", hashMap.Get(99))
	fmt.Printf("227:%v\n", hashMap.Get(227))
	fmt.Printf("355:%v\n", hashMap.Get(355))
	hashMap.Remove(227)
	fmt.Printf("227:%v\n", hashMap.Get(227))
	hashMap.Clear()
	fmt.Printf("137:%v\n", hashMap.Get(137))
}
*/
package hash

import (
	"fmt"
	"testing"
)

func Test_linked_hash_map(t *testing.T) {
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
	hashMap.Remove(9)
	fmt.Println("Remove 9...")
	fmt.Printf("9:%v\n", hashMap.Get(9))
	hashMap.Put(9, "Android")
	fmt.Println("Put 9...")
	fmt.Printf("9:%v\n", hashMap.Get(9))
	hashMap.Remove(227)
	fmt.Println("Remove 227...")
	fmt.Printf("227:%v\n", hashMap.Get(227))
	hashMap.Put(227, "Symbian")
	fmt.Println("Put 227...")
	fmt.Printf("227:%v\n", hashMap.Get(227))
	hashMap.Clear()
	fmt.Println("Clear...")
	fmt.Printf("137:%v\n", hashMap.Get(137))
	hashMap.Put(137, "iOS")
	fmt.Println("Put 137...")
	fmt.Printf("137:%v\n", hashMap.Get(137))
}
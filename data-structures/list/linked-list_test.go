package list

import (
	"fmt"
	"testing"
)

func Test_linked_list(t *testing.T) {
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
	for i := 0; i < list.Size(); i++ {
        	fmt.Printf("%d : %s\n", i, list.Get(i))
	}
	fmt.Printf("Index of 'E': %d\n", list.IndexOf("E"))
	fmt.Printf("First: %s\n", list.GetFirst())
	fmt.Printf("Last: %s\n", list.GetLast())
	fmt.Println("Reversing...")
	list.Reverse()
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("%d : %s\n", i, list.Get(i))
	}
	fmt.Printf("Index of 'E': %d\n", list.IndexOf("E"))
	fmt.Printf("First: %s\n", list.GetFirst())
	fmt.Printf("Last: %s\n", list.GetLast())
}
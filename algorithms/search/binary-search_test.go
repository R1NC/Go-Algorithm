package search

import (
	"fmt"
	"testing"
	)

func Test_binary_search(t *testing.T) {
	sorted_array := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sorted_array)
	x := 11
	fmt.Printf("RecursionBinarySearch '%d': %d\n", x, RecursionBinarySearch(sorted_array, x))
	fmt.Printf("NonRecursionBinarySearch '%d': %d\n", x, NonRecursionBinarySearch(sorted_array, x))
}
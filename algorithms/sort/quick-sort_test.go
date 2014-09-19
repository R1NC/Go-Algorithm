package sort

import (
	"fmt"
	"testing"
)

func Test_quick_sort(t *testing.T) {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	QuickSort(array)
	fmt.Println("QuickSort:", array)
}
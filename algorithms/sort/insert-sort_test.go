package sort

import (
	"fmt"
	"testing"
)

func Test_insert_sort(t *testing.T) {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	InsertSort(array)
	fmt.Println("InsertSort:", array)
}
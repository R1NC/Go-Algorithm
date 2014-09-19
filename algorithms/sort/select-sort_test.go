package sort

import (
	"fmt"
	"testing"
)

func Test_select_sort(t *testing.T) {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	SelectSort(array)
	fmt.Println("SelectSort:", array)
}
package sort

import (
	"fmt"
	"testing"	
)

func Test_bubble_sort(t *testing.T) {
	array1 := []int{2, 1, 3, 5, 6, 4}
	array2 := []int{2, 1, 3, 5, 6, 4}
	array3 := []int{2, 1, 3, 5, 6, 4}
	fmt.Println("Initial array:", array1)
	SimpleBubbleSort(array1)
	fmt.Println("SimpleBubbleSort:", array1)
	FlagSwapBubbleSort(array2)
	fmt.Println("FlagSwapBubbleSort:", array2)
	FlagSwapPositionBubbleSort(array3)
	fmt.Println("FlagSwapPositionBubbleSort:", array3)
}
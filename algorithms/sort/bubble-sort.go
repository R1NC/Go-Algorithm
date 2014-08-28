package main

import "fmt"

func checkArray(array []int) {
	if array == nil {
		panic("Nil array.")
	} else if (len(array) <= 0) {
		panic("Empty array.")
	}
}

func swap(array []int, index1 int, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}

func SimpleBubbleSort(array []int) {
	checkArray(array)
	//var compare int
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array) - i; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
			}
			//compare++
		}
	}
	//fmt.Printf("compare: %d\n", compare)
}

func FlagSwapBubbleSort(array []int) {
	checkArray(array)
	//var compare int
	var has_swapped bool
	for i := 0; i < len(array); i++ {
		has_swapped = false
		for j := 1; j < len(array) - i; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
				has_swapped = true
			}
			//compare++
		}
		// if the last scanning has no swapping, the array is sorted.
		// Then there's no need to scan again.
		if !has_swapped {
			break
		}
	}
	//fmt.Printf("compare: %d\n", compare)
}

func FlagSwapPositionBubbleSort(array []int) {
	checkArray(array)
	//var compare int
	var has_swapped bool
	var flag int
	last_swap_position := len(array)
	for i := 0; i < len(array); i++ {
		// After the last swapping at x, [x, n] is sorted.
		// So we just need to sort [0, x] in next scanning
		flag = last_swap_position
		has_swapped = false
		for j := 1; j < flag; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
				has_swapped = true
				last_swap_position = j
			}
			//compare++
		}
		if !has_swapped {
			break
		}
	}
	//fmt.Printf("compare: %d\n", compare)
}

func main() {
	array1 := []int{2, 1, 3, 5, 6, 4}
	array2 := []int{2, 1, 3, 5, 6, 4}
	array3 := []int{2, 1, 3, 5, 6, 4}
	fmt.Println("Initial array:")
	for _, x := range array1 {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("SimpleBubbleSort...")
	SimpleBubbleSort(array1)
	for _, x := range array1 {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("FlagSwapBubbleSort...")
	FlagSwapBubbleSort(array2)
	for _, x := range array2 {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("FlagSwapPositionBubbleSort...")
	FlagSwapPositionBubbleSort(array3)
	for _, x := range array3 {
		fmt.Printf("%d\n", x)
	}
}

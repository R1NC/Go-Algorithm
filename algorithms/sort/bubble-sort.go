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

func NaiveBubbleSort(array []int) {
	checkArray(array)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swap(array, j, j + 1)
			}
		}
	}
}

func FlagBubbleSort(array []int) {
	checkArray(array)
	var has_swapped bool
	for i := 0; i < len(array); i++ {
		has_swapped = false
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swap(array, j, j + 1)
				has_swapped = true
			}
		}
		// if last scan has not swapped, the array is sorted.
		// Then there's no need to scan again.
		if !has_swapped {
			break
		}
	}
}

func main() {
	array1 := []int{1, 3, 2, 5, 4}
	array2 := []int{1, 3, 2, 5, 4}
	fmt.Println("Initial array:")
	for _, x0 := range array1 {
		fmt.Printf("%d\n", x0)
	}
	fmt.Println("NaiveBubbleSort...")
	NaiveBubbleSort(array1)
	for _, x1 := range array1 {
		fmt.Printf("%d\n", x1)
	}
	fmt.Println("FlagBubbleSort...")
	FlagBubbleSort(array2)
	for _, x2 := range array2 {
		fmt.Printf("%d\n", x2)
	}
}

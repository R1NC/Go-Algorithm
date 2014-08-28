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

func NaiveBubbleSort(array []int) int {
	checkArray(array)
	var compare int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swap(array, j, j + 1)
			}
			compare++
		}
	}
	return compare
}

func FlagBubbleSort(array []int) int {
	checkArray(array)
	var compare int
	var is_last_scan_swapped bool
	for i := 0; i < len(array); i++ {
		is_swapped = false
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swap(array, j, j + 1)
				is_swapped = true
			}
			compare++
		}
		// if last scan has no swapping, the array is sorted.
		// Then there's no need to scan again.
		if !is_swapped {
			break
		}
	}
	return compare
}

func main() {
	array0 := []int{21, 5, 342, 55, 9, 17, 105}
	array1 := array0[:]
	array2 := array0[:]
	//array3 := array0[:]
	fmt.Println("Initial array:")
	for _, x0 := range array0 {
		fmt.Printf("%d\n", x0)
	}
	fmt.Printf("NaiveBubbleSorted with compare: %d\n", NaiveBubbleSort(array1))
	for _, x1 := range array1 {
		fmt.Printf("%d\n", x1)
	}
	fmt.Printf("FlagBubbleSorted with compare: %d\n", FlagBubbleSort(array2))
	for _, x2 := range array2 {
		fmt.Printf("%d\n", x2)
	}
}

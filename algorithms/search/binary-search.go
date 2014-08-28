package main

import "fmt"

func recursion_binary_search(sorted_array []int, target int) int {
	low := 0
	high := len(sorted_array) - 1
	if high < low {
		panic("Empty array.")
	}
	return search(sorted_array, target, low, high)
}

func search(sorted_array []int, target int, low int, high int) int {
	middle := (low + high) / 2
	if sorted_array[middle] == target {
		return middle
	} else if sorted_array[middle] < target {
		return search(sorted_array, target, middle + 1, high)
	} else {
		return search(sorted_array, target, low, middle - 1)
	}
}

func non_recursion_binary_search(sorted_array []int, target int) int {
	low := 0
	high := len(sorted_array) - 1
	if high < low {
		panic("Empty array.")
	}
	for low <= high {
		middle := (low + high) / 2
		if sorted_array[middle] == target {
			return middle
		} else if sorted_array[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

func main() {
	sorted_array := []int{1, 3, 5, 7, 9, 11, 13, 15}
	x := 11
	for _, x := range sorted_array {
		fmt.Printf("%d ", x)
	}
	fmt.Printf("\nrecursion_binary_search '%d': %d\n", x, recursion_binary_search(sorted_array, x))
	fmt.Printf("non_recursion_binary_search '%d': %d\n", x, non_recursion_binary_search(sorted_array, x))
}

package main

import "fmt"

func recursion_binary_search(sorted_array []int, target int) int {
	left := 0
	right := len(sorted_array) - 1
	if right < left {
		panic("Empty array.")
	}
	return search(sorted_array, atrget, left, right)
}

func search(sorted_array []int, target int, left int, right int) int {
	middle := (left + right) / 2
	if sorted_array[middle] == target {
		return middle
	} else if sorted_array[middle] < target {
		search(middle + 1, right)
	} else {
		search(left, middle - 1)
	}
}

func non_recursion_binary_search(sorted_array []int, target int) int {
	left := 0
	right := len(sorted_array) - 1
	if right < left {
		panic("Empty array.")
	}
	for left <= right {
		middle := (left + right) / 2
		if sorted_array[middle] == target {
			return middle
		} else if sorted_array[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	sorted_array := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	x := 11
	for x := range sorted_array {
		fmt.Printf("%d ", x)
	}
	fmt.Printf("\n recursion_binary_search '%d\n': %d", x, recursion_binary_search(x))
	fmt.Printf("non_recursion_binary_search '%d': %d\n", x, non_recursion_binary_search(x))
}
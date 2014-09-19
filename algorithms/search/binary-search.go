/*
  Time complexity: O(logN)
*/

package search

func initSearch(sorted_array []int) (int, int) {
	if sorted_array == nil {
		panic("Nil array.")
	}
	low := 0
	high := len(sorted_array) - 1
	if high < low {
		panic("Empty array.")
	}
	return low, high
}

func NonRecursionBinarySearch(sorted_array []int, target int) int {
	low, high := initSearch(sorted_array)
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

func RecursionBinarySearch(sorted_array []int, target int) int {
	low, high := initSearch(sorted_array)
	return recursion_search(sorted_array, target, low, high)
}

func recursion_search(sorted_array []int, target int, low int, high int) int {
	middle := (low + high) / 2
	if sorted_array[middle] == target {
		return middle
	} else if sorted_array[middle] < target {
		return recursion_search(sorted_array, target, middle + 1, high)
	} else {
		return recursion_search(sorted_array, target, low, middle - 1)
	}
}
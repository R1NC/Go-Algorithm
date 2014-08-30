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

/*
func main() {
	sorted_array := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Array:", sorted_array)
	x := 11
	fmt.Printf("RecursionBinarySearch '%d': %d\n", x, RecursionBinarySearch(sorted_array, x))
	fmt.Printf("NonRecursionBinarySearch '%d': %d\n", x, NonRecursionBinarySearch(sorted_array, x))
}
*/
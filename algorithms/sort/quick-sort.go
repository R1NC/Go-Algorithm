/*
  Time complexity: O(N*logN)
*/
package sort

func partition(array []int, left int, right int) int{
	low := left
	high := right
	pivot := array[(left + right) / 2]
	for low <= high {
		for array[low] < pivot {
			low++
		}
		for array[high] > pivot {
			high--
		}
		if low <= high {
			swap(array, low, high)
			low++
			high--
		}
	}
	return low
}

func recursion(array []int, left int, right int) {
	index := partition(array, left, right)
	if left < index - 1 {
		recursion(array, left, index - 1)
	}
	if index < right {
		recursion(array, index, right)
	}
}

func QuickSort(array []int) {
	checkArray(array)
	recursion(array, 0, len(array) - 1)
}

/*
func main() {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	QuickSort(array)
	fmt.Println("QuickSort:", array)
}
*/
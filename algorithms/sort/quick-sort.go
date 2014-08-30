/*
  QuickSort is suitable for sorting big data.
  The divide-and-conquer strategy is used in QuickSort.
  Time complexity: O(N*logN)
*/
package sort

func partitionRecursion(array []int, left int, right int) {
	low := left
	high := right
	// Choose the middle element as pivot value, but it can be any one.
	pivot := array[(left + right) / 2]
	for low <= high {
		// In the left part of the array, looking for the element which is greater than the pivot.
		for array[low] < pivot {
			low++
		}
		// In the right part of the array, looking for the element which is lesser than the pivot.
		for array[high] > pivot {
			high--
		}
		if low <= high {
			// Swap the two elements we found to meke sure that,
			// all values before 'low' element are less or equal than the pivot,
			// and all values after 'high' element are greater or equal to the pivot.
			swap(array, low, high)
			// Prepare for next scanning.
			low++
			high--
		}
	}
	// Handle the left part of the array which is unsorted.
	if left < high {
		partitionRecursion(array, left, high)
	}
	// Handle the right part of the array which is unsorted.
	if low < right {
		partitionRecursion(array, low, right)
	}
}

func QuickSort(array []int) {
	checkArray(array)
	partitionRecursion(array, 0, len(array) - 1)
}

/*
func main() {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	QuickSort(array)
	fmt.Println("QuickSort:", array)
}
*/
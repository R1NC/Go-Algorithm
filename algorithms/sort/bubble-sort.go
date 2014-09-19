/*
  Time complexity: O(N^2)
*/
package sort

func SimpleBubbleSort(array []int) {
	checkArray(array)
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array) - i; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
			}
		}
	}
}

func FlagSwapBubbleSort(array []int) {
	checkArray(array)
	var has_swapped bool
	for i := 0; i < len(array); i++ {
		has_swapped = false
		for j := 1; j < len(array) - i; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
				has_swapped = true
			}
		}
		// if the last scanning has no swapping, the array is sorted.
		// Then there's no need to scan again.
		if !has_swapped {
			break
		}
	}
}

func FlagSwapPositionBubbleSort(array []int) {
	checkArray(array)
	var has_swapped bool
	var flag int
	last_swap_position := len(array)
	for i := 0; i < len(array); i++ {
		// After the last swapping at x, [x, n] is sorted.
		// So we just need to sort [0, x] in the next scanning.
		flag = last_swap_position
		has_swapped = false
		for j := 1; j < flag; j++ {
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
				has_swapped = true
				last_swap_position = j
			}
		}
		if !has_swapped {
			break
		}
	}
}
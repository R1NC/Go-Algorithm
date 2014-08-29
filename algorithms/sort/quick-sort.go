/*
  Time complexity: O(N*logN)
*/
package main

import "fmt"

func checkArray(array []int) {
	if array == nil {
		panic("Nil array.")
	} else if (len(array) <= 0) {
		panic("Empty array.")
	}
}

/*
func swap(array []int, index1 int, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}
*/

func recursionSort(array []int, left int, right int) {
	key := array[0]
	low := left
	high := right
	for low < high {
		for ; high > low; high-- {
			if array[high] < key {
				array[low] = array[high]
				break
			}
		}
		for ; low < high; low++ {
			if array[low] > key {
				array[high] = array[low]
				break
			}
		}
		array[low] = key
		recursionSort(array, left, low - 1)
		recursionSort(array, low + 1, right)
	}
}

func QuickSort(array []int) {
	checkArray()
	recursionSort(array, 0, len(array) - 1)
}

func main() {
	array := []int{2, 1, 3, 5, 6, 4}
	fmt.Println("Initial array:")
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("QuickSort...")
	QuickSort(array)
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
}
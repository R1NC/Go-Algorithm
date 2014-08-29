/*
  Time complexity: O(N^2)
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

func swap(array []int, index1 int, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}

func InsertSort(array []int) {
	checkArray(array)
	// The array is divided into two parts: sorted-part and unsorted-part. 
	// At the beginning, sorted-part conteins the first element, and unsorted-part contains the rest.
	for i := 1; i < len(array); i++ {
		// For every element in unsorted-part,
		for j := i; j > 0; j-- {
			// choose the right place in sorted-part to insert.
			if array[j - 1] > array[j] {
				swap(array, j - 1, j)
			}
		}
	}
}

func main() {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:")
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("InsertSort...")
	InsertSort(array)
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
}

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

func SelectSort(array []int) {
	checkArray(array)
	// The array is divided into two parts: sorted-part and unsorted-part. 
	// At the beginning, sorted-part is empty, and unsorted-part contains all elements.
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			// Find the minimal element in unsorted-part and add it to the end of sorted-part. 
			if array[j] < array[minIndex] {
				minIndex = j;
			}       
            if minIndex != i {
                  swap(array, i, minIndex)
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
	fmt.Println("SelectSort...")
	SelectSort(array)
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
}

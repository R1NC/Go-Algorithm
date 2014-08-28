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

func NaiveBubbleSort(array []int) {
	checkArray(array)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				swap(array, j, j + 1)
			}
		}
	}
}

func ImprovedBubbleSort(array []int) {
	checkArray(array)
	//TODO
}

func main() {
	array := []int{21, 5, 342, 55, 9, 17, 105}
	fmt.Println("Initial array:")
	for _, x := range array {
		fmt.Printf("%d\n", x)
	}
	NaiveBubbleSort(array)
	fmt.Println("NaiveSorted array:")
	for _, y := range array {
		fmt.Printf("%d\n", y)
	}
}

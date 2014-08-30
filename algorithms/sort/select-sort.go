/*
  Time complexity: O(N^2)
*/
package sort

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

/*
func main() {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	SelectSort(array)
	fmt.Println("SelectSort:", array)
}
*/
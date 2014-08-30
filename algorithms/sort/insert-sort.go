/*
  Time complexity: O(N^2)
*/
package sort

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

/*
func main() {
	array := []int{5, 3, 2, 6, 4, 1}
	fmt.Println("Initial array:", array)
	InsertSort(array)
	fmt.Println("InsertSort:", array)
}
*/
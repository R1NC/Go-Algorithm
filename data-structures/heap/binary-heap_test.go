package heap

import (
	"fmt"
	"testing"
)

func Test_binary_heap(t *testing.T) {
	heap := &BinaryHeap{}
	dataSet := []int{128, 21, 8, 450, 65, 89, 507, 45, 19, 395, 37, 53, 236, 71, 99}
	for _, data := range dataSet {
		heap.Add(data)
	}
	for heap.Size() > 0 {
		fmt.Println(heap.RemoveMinimum())
	}
}
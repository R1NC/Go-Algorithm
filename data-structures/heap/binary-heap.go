package heap

type BinaryHeap struct {
	array []int
}

func (heap *BinaryHeap) Size() int {
	return len(heap.array) - 1
}

func (heap *BinaryHeap) Add(data int) {
	if len(heap.array) == 0 {
		heap.array = append(heap.array, -1)
	}
	heap.array = append(heap.array, data)// Add to the last.
	for i := heap.Size(); heap.array[i] < heap.array[i / 2]; i /= 2 {// Compare with parent.
		tmp := heap.array[i]
		heap.array[i] = heap.array[i / 2]
		heap.array[i / 2] = tmp
	}
}

func (heap *BinaryHeap) RemoveMinimum() int {
	if heap.Size() == 0 {
		panic("Empty heap.")
	}
	minimum := heap.array[1]
	last := heap.array[heap.Size()]
	var i int
	var child int
	for i = 1; i * 2 <= heap.Size(); i = child {
		child = i * 2// Go to left-child.
		if child < heap.Size() && heap.array[child] > heap.array[child + 1] {//Left-child is greater than right-child.
			child++// Go to right-child.
		}
		// Search a 'hole' to store the last element.
		if last > heap.array[child] {//Put the 'hole' to next level.
			tmp := heap.array[i]
			heap.array[i] = heap.array[child]
			heap.array[child] = tmp
		} else {//Found the 'hole'.
			break
		}
	}
	heap.array[i] = last// Store the last element.
	newArray := []int{-1}// Remove the last element.
	for i := 1; i < heap.Size(); i++ {
		newArray = append(newArray, heap.array[i])
	}
	heap.array = newArray
	return minimum
}

/*
func main() {
	heap := &BinaryHeap{}
	dataSet := []int{128, 21, 8, 450, 65, 89, 507, 45, 19, 395, 37, 53, 236, 71, 99}
	for _, data := range dataSet {
		heap.Add(data)
	}
	for heap.Size() > 0 {
		fmt.Println(heap.RemoveMinimum())
	}
}
*/
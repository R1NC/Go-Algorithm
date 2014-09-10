package main

import (
	"fmt"
)

type BinaryHeap struct {
	array []int
}

func (heap *BinaryHeap) Build(dataSet int[]) {
	//TODO
}

func (heap *BinaryHeap) Size() int {
	return len(heap.array)
}

func (heap *BinaryHeap) Add(data int) {
	//TODO
}

func (heap *BinaryHeap) RemoveMinimum() int {
	//TODO
}

func mian() {
	heap := &BinaryHeap{}
	dataSet := []int{21, 65, 89, 45, 19, 37, 53, 71, 99}
	heap.Build(dataSet)
	heap.Add(8)
	for heap.Size() > 0 {
		fmt.Println(heap.RemoveMinimum())
	}
}

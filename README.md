Data Structures
======================
* [LinkedList](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/list/linked-list.go):
```go
func (list *LinkedList) Size() int
```
```go
func (list *LinkedList) Reverse()
```
```go
func (list *LinkedList) IndexOf(value interface{}) int
```
```go
func (list *LinkedList) Get(index int) interface{}
```
```go
func (list *LinkedList) GetFirst() interface{}
```
```go
func (list *LinkedList) GetLast() interface{}
```
```go
func (list *LinkedList) Add(value interface{}, index int)
```
```go
func (list *LinkedList) AddToFirst(value interface{})
```
```go
func (list *LinkedList) AddToLast(value interface{})
```
```go
func (list *LinkedList) RemoveAt(index int)
```
```go
func (list *LinkedList) RemoveFirst()
```
```go
func (list *LinkedList) RemoveLast()
```
* [LinkedStack](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/stack/linked-stack.go):
```go
func (stack *LinkedStack) Size() int
```
```go
func (stack *LinkedStack) Push(value interface{})
```
```go
func (stack *LinkedStack) Pop()
```
```go
func (stack *LinkedStack) Peek() interface{}
```
* [LinkedQueue](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/queue/linked-queue.go):
```go
func (queue *LinkedQueue) Size() int
```
```go
func (queue *LinkedQueue) Add(value interface{})
```
```go
func (queue *LinkedQueue) Remove()
```
```go
func (queue *LinkedQueue) Peek() interface{}
```
* [LinkedHashMap](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/hash/linked-hash-map.go):
```go
func (hashMap *LinkedHashMap) Put(key int, value interface{})
```
```go
func (hashMap *LinkedHashMap) Get(key int) interface{}
```
```go
func (hashMap *LinkedHashMap) Remove(key int)
```
```go
func (hashMap *LinkedHashMap) Clear()
```
* [BinarySearchTree](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/tree/binary-search-tree.go)
```go
func (tree *BinarySearchTree) Add(value int)
```
```go
func (tree *BinarySearchTree) Remove(value int)
```
```go
func (tree *BinarySearchTree) Search(value int) *BinarySearchTree
```
```go
func (tree *BinarySearchTree) Traverse()
```
```go
func (tree *BinarySearchTree) TraverseByLevel()
```
* [BinaryHeap](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/heap/binary-heap.go)
```go
func (heap *BinaryHeap) Size() int
```
```go
func (heap *BinaryHeap) Add(data int)
```
```go
func (heap *BinaryHeap) RemoveMinimum() int
```
* [Graph](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/graph/graph.go):
```go
func (graph *Graph) BreadFirstTraverse(startVertex *Vertex)
```
```go
func (graph *Graph) DepthFirstTraverse(startVertex *Vertex)
```
```go
func (graph *Graph) PrimMinimumSpanningTree(startVertex *Vertex)
```
```go
func (graph *Graph) KruskalMinimumSpanningTree()
```
```go
func (graph *Graph) DijkstraShortestPath(startVertex *Vertex, endVertex *Vertex)
```
```go
func (graph *Graph) TopologicalSort()
```
Sorting Algorithms
===============
* [BubbleSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/bubble-sort.go):
```go
func SimpleBubbleSort(array []int)
```
```go
func FlagSwapBubbleSort(array []int)
```
```go
func FlagSwapPositionBubbleSort(array []int)
```
* [InsertSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/insert-sort.go):
```go
func InsertSort(array []int)
```
* [SelectSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/select-sort.go):
```go
func SelectSort(array []int)
```
* [QuickSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/quick-sort.go):
```go
func QucikSort(array []int)
```
Searching Algorithms
=================
* [BinarySearch](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/search/binary-search.go):
```go
func RecursionBinarySearch(sorted_array []int, target int) int
```
```go
func NonRecursionBinarySearch(sorted_array []int, target int) int
```

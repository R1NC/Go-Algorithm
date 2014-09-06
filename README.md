Linear Data Structures
======================
* [LinkedList](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/list/linked-list.go):
```go
func (list *LinkedList) Size() int
func (list *LinkedList) Reverse()
func (list *LinkedList) IndexOf(value interface{}) int
func (list *LinkedList) Get(index int) interface{}
func (list *LinkedList) GetFirst() interface{}
func (list *LinkedList) GetLast() interface{}
func (list *LinkedList) Add(value interface{}, index int)
func (list *LinkedList) AddToFirst(value interface{})
func (list *LinkedList) AddToLast(value interface{})
func (list *LinkedList) RemoveAt(index int)
func (list *LinkedList) RemoveFirst()
func (list *LinkedList) RemoveLast()
```
* [LinkedStack](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/stack/linked-stack.go):
```go
func (stack *LinkedStack) Size() int
func (stack *LinkedStack) Push(value interface{})
func (stack *LinkedStack) Pop()
func (stack *LinkedStack) Peek() interface{}
```
* [LinkedQueue](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/queue/linked-queue.go):
```go
func (queue *LinkedQueue) Size() int
func (queue *LinkedQueue) Add(value interface{})
func (queue *LinkedQueue) Remove()
func (queue *LinkedQueue) Peek() interface{}
```
* [LinkedHashMap](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/hash/linked-hash-map.go):
```go
func (hashMap *LinkedHashMap) Put(key int, value interface{}) 
func (hashMap *LinkedHashMap) Get(key int) interface{}
func (hashMap *LinkedHashMap) Remove(key int)
func (hashMap *LinkedHashMap) Clear()
```
Nonlinear Data Structures
========================
* [BinaryTree](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/tree/binary-tree.go):
```go
func (tree *BinaryTree) Add(value int)
func (tree *BinaryTree) Remove(value int)
func (tree *BinaryTree) Search(value int) *BinaryTree
func (tree *BinaryTree) Traverse()
func (tree *BinaryTree) TraverseByLevel()
```
* [Graph](https://github.com/RincLiu/Go-Algorithm/blob/master/data-structures/graph/graph.go):
```go
func (graph *Graph) BreadFirstTraverse(startVertex *Vertex)
func (graph *Graph) DepthFirstTraverse(startVertex *Vertex)
func (graph *Graph) PrimMinimumSpanningTree(startVertex *Vertex)
func (graph *Graph) KruskalMinimumSpanningTree()
func (graph *Graph) DijkstraShortestPath(startVertex *Vertex, endVertex *Vertex)
func (graph *Graph) TopologicalSort()
```
Sorting Algorithms
===============
* [BubbleSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/bubble-sort.go):
```go
func SimpleBubbleSort(array []int)
func FlagSwapBubbleSort(array []int)
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
func NonRecursionBinarySearch(sorted_array []int, target int) int
```

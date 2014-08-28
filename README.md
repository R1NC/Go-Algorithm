Data Structures
===============
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
Sort Algorithms
===============
* [BubbleSort](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/sort/bubble-sort.go):
```go
func NaiveBubbleSort(array []int)
func ImprovedBubbleSort(array []int)
```
Search Algorithms
=================
* [BinarySearch](https://github.com/RincLiu/Go-Algorithm/blob/master/algorithms/search/binary-search.go):
```go
func RecursionBinarySearch(sorted_array []int, target int) int
func NonRecursionBinarySearch(sorted_array []int, target int) int
```
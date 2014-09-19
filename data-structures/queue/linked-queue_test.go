package queue

import (
	"fmt"
	"testing"
)

func Test_linked_queue(t *testing.T) {
	queue := &LinkedQueue{}
	queue.Add("A")
	queue.Add("B")
	queue.Add("C")
	queue.Add("D")
	for queue.Size() > 0 {
		fmt.Printf("%s\n", queue.Peek())
		queue.Remove()
	}
}
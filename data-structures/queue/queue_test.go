package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{0, nil, nil}

	// push
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	fmt.Println(q.Size())
	fmt.Println("------------------------------")

	// pop
	node := q.Pop()
	fmt.Println(node)
	fmt.Println(q.Size())
	node = q.Pop()
	fmt.Println(node)
	fmt.Println(q.Size())
}

package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{nil, 0}

	// push
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println(s.Size())
	fmt.Println("------------------------------")

	// pop
	node := s.Pop()
	fmt.Println(node)
	fmt.Println(s.Size())
	node = s.Pop()
	fmt.Println(node)
	fmt.Println(s.Size())
}

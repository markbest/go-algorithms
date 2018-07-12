package list

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	l := DoubleList{0, nil, nil}

	// append
	for i := 0; i < 10; i++ {
		l.Append(i)
	}
	l.Traverse()
	fmt.Println("------------------------------")

	// insert
	l.Insert(11, 8)
	l.Traverse()
	fmt.Println("------------------------------")

	// remove
	l.Remove(5)
	l.Traverse()
	fmt.Println("------------------------------")

	// get
	node := l.Get(3)
	fmt.Println(node)
}

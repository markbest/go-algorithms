package list

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	l := &SingleList{0, nil, nil}

	// append node
	for i := 0; i <= 10; i++ {
		l.Append(&SingleNode{i, nil})
	}
	l.Traverse()
	fmt.Println("------------------------------")

	// insert node
	l.Insert(&SingleNode{14, nil}, 4)
	l.Traverse()
	fmt.Println("------------------------------")

	// remove node
	l.Remove(5)
	l.Traverse()
	fmt.Println("------------------------------")

	// get node
	node := l.Get(4)
	fmt.Println(node)
}

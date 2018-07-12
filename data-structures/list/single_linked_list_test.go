package list

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	n := &LinkNode{}

	for i := 0; i <= 10; i++ {
		n.Insert(i, i)
	}
	n.Traverse()
	fmt.Printf("first node is %v\n", n.First())
	fmt.Printf("last node is %v\n", n.Last())
	fmt.Printf("length of list is %d\n", n.Length())

	n.Delete(5)
	n.Traverse()
	fmt.Printf("first node is %v\n", n.First())
	fmt.Printf("last node is %v\n", n.Last())
	fmt.Printf("length of list is %d\n", n.Length())
}

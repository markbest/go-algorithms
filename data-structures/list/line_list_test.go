package list

import (
	"testing"
)

func TestLineList(t *testing.T) {
	list := LineList{make([]int, 0), 0}

	// append
	for i := 0; i <= 10; i++ {
		list.Append(i)
	}
	list.Traverse()

	// insert
	list.Insert(18, 8)
	list.Traverse()

	// remove
	list.Remove(5)
	list.Traverse()
}

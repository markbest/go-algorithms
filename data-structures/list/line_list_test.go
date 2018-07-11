package list

import (
	"testing"
)

func TestLineList(t *testing.T) {
	list := NewLineList()
	list.Insert(18, 8)
	t.Log(list.Dump())
	t.Log(list.Len())

	list.Delete(8)
	t.Log(list.Dump())
	t.Log(list.Len())
}

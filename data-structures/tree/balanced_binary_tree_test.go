package tree

import (
	"testing"
)

func TestAvlTree(t *testing.T) {
	nodes := []int{4, 3, 2, 7, 9, 11, 10}
	var tr *AVLTree
	for _, v := range nodes {
		tr = Insert(tr, v)
	}
	Traverse(tr)
}

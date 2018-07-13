package tree

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

// insert node data
func (n *BinaryNode) Insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

type BinaryTree struct {
	root *BinaryNode
}

// insert tree node
func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{left: nil, right: nil, data: data}
	} else {
		t.root.Insert(data)
	}
	return t
}

// find tree min node
func (t *BinaryTree) FindMin() int {
	node := t.root
	for {
		if node.left != nil {
			node = node.left
		} else {
			return node.data
		}
	}
}

// find tree max node
func (t *BinaryTree) FindMax() int {
	node := t.root
	for {
		if node.right != nil {
			node = node.right
		} else {
			return node.data
		}
	}
}

// find tree node
func (t *BinaryTree) Find(v int) bool {
	node := t.root
	for {
		if node == nil {
			return false
		} else if node.data == v {
			return true
		} else if node.data > v {
			node = node.left
		} else {
			node = node.right
		}
	}
}

// print tree node
func (t *BinaryTree) Print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	t.Print(w, node.left, ns+2, 'L')
	t.Print(w, node.right, ns+2, 'R')
}

package tree

import (
	"fmt"
)

type BinarySearchTree struct {
	root *BinarySearchNode
}

type BinarySearchNode struct {
	parent *BinarySearchNode
	left   *BinarySearchNode
	right  *BinarySearchNode
	data   int
}

// insert node data
func (n *BinarySearchNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinarySearchNode{left: nil, right: nil, parent: n, data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinarySearchNode{left: nil, right: nil, parent: n, data: data}
		} else {
			n.right.insert(data)
		}
	}
}

// insert tree node
func (t *BinarySearchTree) Insert(data int) *BinarySearchTree {
	if t.root == nil {
		t.root = &BinarySearchNode{left: nil, right: nil, parent: nil, data: data}
	} else {
		t.root.insert(data)
	}
	return t
}

// find tree min node
func (t *BinarySearchTree) FindMin() int {
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
func (t *BinarySearchTree) FindMax() int {
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
func (t *BinarySearchTree) Contain(v int) bool {
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

// remove tree node
func (t *BinarySearchTree) Remove(v int) bool {
	node := t.root
	for {
		if node == nil {
			return false
		} else if node.data == v {
			if node.left != nil && node.right != nil {
				tr := &BinarySearchTree{root: node.right}
				min := tr.FindMin()
				t.Remove(min)
				node.data = min
			}
			if node.left != nil && node.right == nil {
				node.parent.left = node.left
			}
			if node.left == nil && node.right != nil {
				node.parent.right = node.right
			}
			if node.left == nil && node.right == nil {
				if node.parent.left != nil && node.parent.left.data == v {
					node.parent.left = nil
				}
				if node.parent.right != nil && node.parent.right.data == v {
					node.parent.right = nil
				}
			}
			return true
		} else if node.data > v {
			node = node.left
		} else {
			node = node.right
		}
	}
}

// pre order traverse
func (t *BinarySearchTree) PreTraverse(node *BinarySearchNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.data)
	t.PreTraverse(node.left)
	t.PreTraverse(node.right)
}

// mid order traverse
func (t *BinarySearchTree) MidTraverse(node *BinarySearchNode) {
	if node == nil {
		return
	}

	t.MidTraverse(node.left)
	fmt.Printf("%d ", node.data)
	t.MidTraverse(node.right)
}

// next order traverse
func (t *BinarySearchTree) NextTraverse(node *BinarySearchNode) {
	if node == nil {
		return
	}

	t.NextTraverse(node.left)
	t.NextTraverse(node.right)
	fmt.Printf("%d ", node.data)
}

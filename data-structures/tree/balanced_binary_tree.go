package tree

import (
	"fmt"
)

type AVLTree struct {
	data  int
	high  int
	left  *AVLTree
	right *AVLTree
}

// get tree high
func getTreeHigh(t *AVLTree) int {
	if t == nil {
		return -1
	} else {
		return t.high
	}
}

// get max value
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// insert tree data
func Insert(avl *AVLTree, data int) *AVLTree {
	if avl == nil {
		avl = &AVLTree{data: data, high: 0, left: nil, right: nil}
		avl.high = max(getTreeHigh(avl.left), getTreeHigh(avl.right)) + 1
		return avl
	}

	if data < avl.data {
		avl.left = Insert(avl.left, data)
		if getTreeHigh(avl.left)-getTreeHigh(avl.right) == 2 {
			if data < avl.left.data {
				avl = RightRotate(avl)
			} else {
				avl = LeftRightRotate(avl)
			}
		}
	}

	if data > avl.data {
		avl.right = Insert(avl.right, data)
		if getTreeHigh(avl.right)-getTreeHigh(avl.left) == 2 {
			if data < avl.right.data {
				avl = RightLeftRotate(avl)
			} else {
				avl = LeftRotate(avl)
			}
		}
	}

	avl.high = max(getTreeHigh(avl.left), getTreeHigh(avl.right)) + 1
	return avl
}

// tree right rotate
func RightRotate(avl *AVLTree) *AVLTree {
	temp := avl.left
	avl.left = temp.right
	temp.right = avl
	avl.high = max(getTreeHigh(avl.left), getTreeHigh(avl.right)) + 1
	temp.high = max(getTreeHigh(temp.left), avl.high) + 1
	return temp
}

// tree right rotate
func LeftRotate(avl *AVLTree) *AVLTree {
	temp := avl.right
	avl.right = temp.left
	temp.left = avl
	avl.high = max(getTreeHigh(avl.left), getTreeHigh(avl.right)) + 1
	temp.high = max(avl.high, getTreeHigh(temp.right)) + 1
	return temp
}

// tree left right rotate
func LeftRightRotate(avl *AVLTree) *AVLTree {
	avl.left = LeftRotate(avl.left)
	avl = RightRotate(avl)
	return avl
}

// tree right left rotate
func RightLeftRotate(avl *AVLTree) *AVLTree {
	avl.right = RightRotate(avl.right)
	avl = LeftRotate(avl)
	return avl
}

// traverse tree
func Traverse(avl *AVLTree) {
	if avl == nil {
		return
	}

	fmt.Printf("%d ", avl.data)
	Traverse(avl.left)
	Traverse(avl.right)
}

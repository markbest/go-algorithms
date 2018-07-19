package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinarySearchTree{}
	elements := []int{100, -20, -50, -15, -60, 50, 60, 55, 85, 15, 5, -10}
	for _, v := range elements {
		tree.Insert(v)
	}
	fmt.Printf("最小节点: %d\n", tree.FindMin())
	fmt.Printf("最大节点: %d\n", tree.FindMax())
	fmt.Printf("查询节点: %d，查询结果: %t\n", 15, tree.Contain(15))
	fmt.Println("前序遍历:")
	tree.PreTraverse(tree.root)
	fmt.Println()
	fmt.Println("中序遍历:")
	tree.MidTraverse(tree.root)
	fmt.Println()
	fmt.Println("后序遍历:")
	tree.NextTraverse(tree.root)
}

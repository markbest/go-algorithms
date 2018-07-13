package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.Insert(100)
	tree.Insert(-20)
	tree.Insert(-50)
	tree.Insert(-15)
	tree.Insert(-60)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(55)
	tree.Insert(85)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(-10)
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

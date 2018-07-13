package tree

import (
	"fmt"
	"os"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)
	tree.Print(os.Stdout, tree.root, 0, 'M')
	fmt.Println()
	fmt.Printf("min node value is: %d\n", tree.FindMin())
	fmt.Printf("max node value is: %d\n", tree.FindMax())
	fmt.Printf("search node value: %d result is: %t", 15, tree.Find(15))
	fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println()

	tree = &BinaryTree{}
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	tree.Print(os.Stdout, tree.root, 0, 'M')
	fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println()

	tree = &BinaryTree{}
	for i := 10; i >= 0; i-- {
		tree.Insert(i)
	}
	tree.Print(os.Stdout, tree.root, 0, 'M')
}

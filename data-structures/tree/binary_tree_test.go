package tree

import (
	"fmt"
	"os"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	tree.print(os.Stdout, tree.root, 0, 'M')
	fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println()

	tree = &BinaryTree{}
	for i := 0; i < 10; i++ {
		tree.insert(int64(i))
	}
	tree.print(os.Stdout, tree.root, 0, 'M')
	fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println()

	tree = &BinaryTree{}
	for i := 10; i >= 0; i-- {
		tree.insert(int64(i))
	}
	tree.print(os.Stdout, tree.root, 0, 'M')
}

package list

import (
	"fmt"
)

type SingleList struct {
	size int
	head *SingleNode
	tail *SingleNode
}

type SingleNode struct {
	data int
	next *SingleNode
}

// append node
func (l *SingleList) Append(e int) bool {
	newNode := &SingleNode{data: e}
	if l.size == 0 {
		l.head = newNode
	} else {
		oldTail := l.tail
		oldTail.next = newNode
	}
	l.tail = newNode
	l.size++
	return true
}

// insert node
func (l *SingleList) Insert(e int, pos int) bool {
	if pos == 0 || pos >= l.size {
		return false
	}

	newNode := &SingleNode{data: e}
	if pos == 1 {
		newNode.next = l.head
		l.head = newNode
	} else {
		preItem := l.head
		for i := 2; i < pos; i++ {
			preItem = preItem.next
		}

		newNode.next = preItem.next
		preItem.next = newNode
	}
	l.size++
	return true
}

// remove node
func (l *SingleList) Remove(pos int) bool {
	if pos >= l.size || pos == 0 {
		return false
	}

	if pos == 1 {
		preItem := l.head
		l.head = preItem.next
	} else {
		preItem := l.head
		for i := 1; i < pos-1; i++ {
			preItem = preItem.next
		}

		node := preItem.next
		preItem.next = node.next

		if pos == l.size {
			l.tail = preItem
		}
	}

	l.size--
	return true
}

// get node
func (l *SingleList) Get(pos int) *SingleNode {
	if pos >= l.size {
		return nil
	}

	item := l.head
	for i := 1; i < pos; i++ {
		item = item.next
	}
	return item
}

// traverse list
func (l *SingleList) Traverse() {
	item := l.head
	fmt.Println(item)
	for i := 0; i < l.size-1; i++ {
		fmt.Println(item.next)
		item = item.next
	}
}

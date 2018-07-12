package list

import (
	"fmt"
)

type DoubleList struct {
	size int
	head *DoubleNode
	tail *DoubleNode
}

type DoubleNode struct {
	data int
	prev *DoubleNode
	next *DoubleNode
}

// append element
func (l *DoubleList) Append(e int) {
	newNode := &DoubleNode{data: e}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
		newNode.prev = nil
		newNode.next = nil
	} else {
		newNode.prev = l.tail
		newNode.next = nil
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// insert element
func (l *DoubleList) Insert(e int, pos int) bool {
	if pos == 0 || pos >= l.size {
		return false
	}

	newNode := &DoubleNode{data: e}
	if pos == 1 {
		preItem := l.head
		newNode.prev = nil
		newNode.next = preItem
		preItem.prev = newNode
		l.head = newNode
	} else {
		preItem := l.head
		for i := 1; i < pos; i++ {
			preItem = preItem.next
		}

		newNode.prev = preItem.prev
		newNode.next = preItem
		preItem.prev.next = newNode
	}
	l.size++
	return true
}

// remove element
func (l *DoubleList) Remove(pos int) bool {
	if pos == 0 || pos >= l.size {
		return false
	}

	if pos == 1 {
		preItem := l.head
		preItem.next.prev = nil
		l.head = preItem.next
	} else {
		preItem := l.head
		for i := 1; i < pos; i++ {
			preItem = preItem.next
		}

		preItem.prev.next = preItem.next
		preItem.next.prev = preItem.prev
	}
	l.size--
	return true
}

// get node
func (l *DoubleList) Get(pos int) *DoubleNode {
	if pos == 0 || pos > l.size {
		return nil
	}

	preItem := l.head
	for i := 1; i < pos; i++ {
		preItem = preItem.next
	}
	return preItem
}

// traverse list
func (l *DoubleList) Traverse() {
	item := l.head
	fmt.Println(item)
	for i := 0; i < l.size-1; i++ {
		fmt.Println(item.next)
		item = item.next
	}
}

package list

import (
	"fmt"
)

type LinkNode struct {
	Data int
	Next *LinkNode
}

// get first element
func (l *LinkNode) First() *LinkNode {
	if l.Next == nil {
		return nil
	}
	return l.Next
}

// get last element
func (l *LinkNode) Last() *LinkNode {
	if l.Next == nil {
		return nil
	}

	point := l.Next
	for point != nil {
		point = point.Next
	}
	return point
}

// insert element
func (l *LinkNode) Insert(e int, i int) bool {
	p := l
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p == nil || j > i {
		return false
	}
	s := &LinkNode{Data: e}
	s.Next = p.Next
	p.Next = s
	return true
}

// delete element
func (l *LinkNode) Delete(i int) bool {
	p := l
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p == nil || j > i {
		return false
	}
	p.Next = p.Next.Next
	return true
}

// get element
func (l *LinkNode) Get(i int) int {
	point := l.Next
	for j := 1; j < i; j++ {
		if point == nil {
			return 0
		}
		point = point.Next
	}
	return point.Data
}

// Traverse all element
func (l *LinkNode) Traverse() {
	point := l.Next
	for point != nil {
		fmt.Println(point.Data, point.Next)
		point = point.Next
	}
}

// get length
func (l *LinkNode) Length() int {
	length := 0
	point := l.Next
	for point != nil {
		length++
		point = point.Next
	}
	return length
}

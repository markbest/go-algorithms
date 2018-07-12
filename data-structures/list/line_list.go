package list

import "fmt"

type LineList struct {
	data []int
	size int
}

// append element
func (l *LineList) Append(e ...int) {
	l.data = append(l.data, e...)
	l.size++
}

// insert element
func (l *LineList) Insert(e int, pos int) {
	if pos > l.size {
		for i := l.size; i < pos-1; i++ {
			l.data = append(l.data, 0)
		}
		l.data = append(l.data, e)
		l.size = pos
	}
	if pos < l.size {
		temp := append([]int{}, l.data[pos-1:]...)
		l.data = append(l.data[:pos-1], e)
		l.data = append(l.data, temp...)
		l.size++
	}
}

// delete element
func (l *LineList) Remove(pos int) {
	if pos < l.size {
		l.data = append(l.data[:pos-1], l.data[pos:]...)
		l.size--
	}
	if pos == l.size {
		l.data = l.data[:pos-1]
		l.size--
	}
}

// dump element
func (l *LineList) Traverse() {
	fmt.Println(l.data)
}

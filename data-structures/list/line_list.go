package list

type lineList struct {
	elements []int
	length   int
}

// new line list
func NewLineList() *lineList {
	return &lineList{elements: make([]int, 0), length: 0}
}

// insert element
func (l *lineList) Insert(e int, pos int) {
	if pos > l.length {
		for i := l.length; i < pos-1; i++ {
			l.elements = append(l.elements, 0)
		}
		l.elements = append(l.elements, e)
		l.length = pos
	}
	if pos < l.length {
		temp := l.elements
		l.elements = temp[:pos-1]
		l.elements = append(l.elements, e)
		l.elements = append(l.elements, temp[pos-1:]...)
		l.length++
	}
}

// delete element
func (l *lineList) Delete(pos int) {
	if pos < l.length {
		temp := l.elements
		l.elements = temp[:pos-1]
		l.elements = append(l.elements, temp[pos:]...)
		l.length--
	}
	if pos == l.length {
		l.elements = l.elements[:pos-1]
		l.length--
	}
}

// get list length
func (l *lineList) Len() int {
	return l.length
}

// dump element
func (l *lineList) Dump() []int {
	return l.elements
}

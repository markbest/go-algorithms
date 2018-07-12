package stack

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	data int
	next *Node
}

// push element
func (s *Stack) Push(e int) {
	newNode := &Node{e, s.top}
	s.top = newNode
	s.size++
}

// pop element
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	temp := s.top
	s.top = temp.next
	s.size--
	return temp
}

// is stack empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// get stack size
func (s *Stack) Size() int {
	return s.size
}

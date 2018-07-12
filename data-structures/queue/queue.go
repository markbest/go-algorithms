package queue

type Queue struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
}

// push element
func (q *Queue) Push(e int) {
	newNode := &Node{e, nil}
	if q.head == nil {
		q.head = newNode
	} else {
		preItem := q.head
		for i := 1; i < q.size; i++ {
			preItem = preItem.next
		}
		preItem.next = newNode
	}
	q.tail = newNode
	q.size++
}

// pop element
func (q *Queue) Pop() *Node {
	if q.IsEmpty() {
		return nil
	}
	temp := q.head
	q.head = temp.next
	q.size--
	return temp
}

// is queue empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// get queue size
func (q *Queue) Size() int {
	return q.size
}

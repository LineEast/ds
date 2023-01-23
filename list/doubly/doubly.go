package doubly

type (
	List[T any] struct {
		head, tail *Node[T]

		len uint
	}
)

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Len() uint {
	return l.len
}

// Clears List and nodes
func (l *List[T]) FullClear() {
	for n := l.head; n != nil; {
		last := n.next
		n.next = nil
		n = last
	}

	l.head, l.tail = nil, nil

	l.len = 0
}

// Clears List
func (l *List[T]) Clear() {
	l.head, l.tail = nil, nil

	l.len = 0
}

func (l *List[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Back() *Node[T] {
	return l.tail
}

func (l *List[T]) PushFront(body T) {
	node := NewNode(body)

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.head
		l.head.prev = node
		l.head = node
	}

	l.len++
}

func (l *List[T]) PushNodeFront(node *Node[T]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.head
		l.head.prev = node
		l.head = node
	}

	l.len++
}

func (l *List[T]) PushListFront(list *List[T]) {
	if l.head == nil {
		l.head = list.head
		l.tail = list.tail
	} else {
		l.tail.next = list.head
		l.tail = list.tail
	}

	l.len += list.len
}

func (l *List[T]) PushBack(body T) {
	node := NewNode(body)

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}

	l.len++
}

func (l *List[T]) PushNodeBack(node *Node[T]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}

	l.len++
}

func (l *List[T]) PopFront() (n T) {
	if l.head == nil {
		return
	}

	n = l.head.body

	l.Remove(l.head)

	return
}

func (l *List[T]) PopBack() (n T) {
	if l.head == nil {
		return
	}

	n = l.tail.body

	l.Remove(l.tail)

	return
}

func (l *List[T]) PopNodeFront() (n *Node[T]) {
	if l.head == nil {
		return nil
	}

	n = l.head

	l.Remove(l.head)

	return n
}

func (l *List[T]) PopNodeBack() (n *Node[T]) {
	if l.IsEmpty() {
		return nil
	}

	n = l.tail

	l.Remove(l.tail)

	return n
}

func (l *List[T]) Remove(n *Node[T]) {
	n.prev.next = n.next
	n.next.next = n.prev

	n.next = nil
	n.prev = nil

	l.len--
}

package singly

type (
	List[T any] struct {
		root, tail *Node[T]

		len uint
	}
)

// New returns an initialized list.
func New[T any]() *List[T] {
	return &List[T]{}
}

// Len returns uint len of the list l
func (l *List[T]) Len() uint {
	return l.len
}

// Clears list l and all nodes links in it.
func (l *List[T]) Clear() {
	for n := l.root; n != nil; {
		last := n.Next()
		n.SetNext(nil)
		n = last
	}

	l.root, l.tail = nil, nil

	l.len = 0
}

// Check if l List is empty
func (l *List[T]) IsEmpty() bool {
	return l.root == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Node[T] {
	return l.root
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Node[T] {
	return l.tail
}

// Push to the end of the l list
func (l *List[T]) PushBack(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.root = node
		l.tail = node
	} else {
		l.Back().next = node
		l.tail = node
	}

	l.len++
}

// Push n node to the end of the l list
func (l *List[T]) PushNodeBack(n *Node[T]) {
	if l.IsEmpty() {
		l.root = n
		l.tail = n
	} else {
		l.Back().next = n
		l.tail = n
	}

	l.len++
}

// Push list List to the back of the l List
func (l *List[T]) PushListBack(list *List[T]) {
	if l.IsEmpty() {
		l.root = list.Front()
		l.tail = list.Back()
	} else {
		l.Back().next = list.Front()
		l.tail = list.Back()
	}

	l.len += list.Len()
}

// Push to the front of the l list
func (l *List[T]) PushFront(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.root = node
		l.tail = node
	} else {
		node.next = l.root
		l.root = node
	}

	l.len++
}

// Push n node to the front of the l list
func (l *List[T]) PushNodeFront(n *Node[T]) {
	if l.IsEmpty() {
		l.root = n
		l.tail = n
	} else {
		n.next = l.root
		l.root = n
	}

	l.len++
}

// Push list List to the front of the l List
func (l *List[T]) PushListFront(list *List[T]) {
	list.PushListBack(l)
	l = list
}

func (l *List[T]) PopFront() (v T) {
	v = l.root.body
	l.Remove(l.Front())
	return
}

func (l *List[T]) PopBack() (v T) {
	v = l.tail.body

	l.Remove(l.Back())

	return v
}

// Remove removes e from l if e is an element of list l
func (l *List[T]) Remove(node *Node[T]) {
	if l.root == nil {
		return
	}

	last := l.root

	if last == node {
		l.root = l.root.next
		node.next = nil

		if l.root == nil {
			l.tail = nil
		}

		l.len--

		return
	}

	for n := l.root; n.next != node || n.next != nil; n = n.next {
		if n == node {
			if n == l.tail {
				l.tail = last
			}

			last.next = n.next
			node.next = nil
			l.len--

			return
		}

		last = n
	}
}

// Remove removes e from l if e is an element of list l
func (l *List[T]) RemoveByBody(compare func(T, T) bool, body T) {
	if l.root == nil {
		return
	}

	last := l.root
	if compare(l.root.body, body) {
		l.root = last.next
		last.next = nil

		if l.root == nil {
			l.tail = nil
		}

		l.len--

		return
	}

	for n := l.root; n.next != nil; n = n.next {
		if compare(n.body, body) {
			if n == l.tail {
				l.tail = last
			}

			last.next = n.next
			n.next = nil

			l.len--

			return
		}

		last = n
	}
}

// func (l *List[T]) RemoveByBody(body T) {
// 	if l.root == nil {
// 		return
// 	}

// 	last := l.root
// 	if last == node {
// 		l.root = l.root.next
// 		node.next = nil

// 		l.len--

// 		return
// 	}

// 	for n := l.root; n.next != node || n.next != nil; n = n.next {
// 		if n == node {
// 			if n == l.tail {
// 				l.tail = last
// 			}

// 			last.next = n.next
// 			node.next = nil
// 			l.len--

// 			return
// 		}

// 		last = n
// 	}
// }

package doubly

type (
	DoublyOrdered[T any] struct {
		list *List[T]

		Compare func(a, b T) bool
	}
)

// Get new ordered list
func NewOrderedList[T any](c func(a, b T) bool) *DoublyOrdered[T] {
	return &DoublyOrdered[T]{
		list:    New[T](),
		Compare: c,
	}
}

// Get len of list
func (l *DoublyOrdered[T]) Len() uint {
	return l.list.len
}

// Push new node into list
//
// Returns node
func (l *DoublyOrdered[T]) Push(body T) (node *Node[T]) {
	node.body = body

	if l.list.head == nil {
		l.list.head = node
		l.list.tail = node
	} else if l.Compare(body, l.list.tail.body) {
		node.prev = l.list.tail
		l.list.tail.next = node
		l.list.tail = node
	} else {
		for n := l.list.head; n != nil; n = n.next {
			if l.Compare(n.body, body) {
				node.prev = n.prev
				node.next = n

				node.prev.next = node
				n.prev = node
			}
		}
	}

	l.list.len++

	return
}

func (l *DoublyOrdered[T]) PushNode(node *Node[T]) {
	if l.list.head == nil {
		l.list.head = node
		l.list.tail = node
	} else if l.Compare(node.body, l.list.tail.body) {
		node.prev = l.list.tail
		l.list.tail.next = node
		l.list.tail = node
	} else {
		for n := l.list.head; n != nil; n = n.next {
			if l.Compare(n.body, node.body) {
				node.prev = n.prev
				node.next = n

				node.prev.next = node
				n.prev = node
			}
		}
	}

	l.list.len++

	return
}

package singly

type (
	SinglyOrdered[T any] struct {
		List *List[T]

		Compare func(a, b T) bool
	}
)

// Get new single ordered linked list
//
// All methods with list from l.List
func NewOrderdList[T any](c func(v1, v2 T) bool) *SinglyOrdered[T] {
	return &SinglyOrdered[T]{
		List:    New[T](),
		Compare: c,
	}
}

// push into single ordered list
func (l *SinglyOrdered[T]) Push(body T) {
	node := &Node[T]{
		body: body,
	}

	if l.List.root == nil {
		l.List.root = node
		l.List.tail = node

		l.List.len++

		return
	}

	last := l.List.root
	for n := l.List.root; n != nil; n = n.next {
		if l.Compare(n.body, body) {
			if n == l.List.root {
				node.next = l.List.root
				l.List.root = node
			} else {
				node.next = last.next
				last.next = node
			}
		} else if n.next == nil {
			l.List.tail = node
			n.next = node
		} else {
			last = n
			continue
		}

		l.List.len++
		return
	}
}

func (l *SinglyOrdered[T]) PushNode(node *Node[T]) {
	if l.List.IsEmpty() {
		l.List.root = node
		l.List.tail = node

		l.List.len++

		return
	}

	last := l.List.root
	for n := l.List.root; n != nil; n = n.next {
		if l.Compare(n.body, node.body) {
			if n == l.List.root {
				node.next = l.List.root
				l.List.root = node
			} else {
				node.next = last.next
				last.next = node
			}
		} else if n.next == nil {
			l.List.tail = node
			n.next = node
		} else {
			last = n
			continue
		}

		l.List.len++
		return
	}
}

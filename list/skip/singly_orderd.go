package skip

type (
	singlyOrdered[T any] struct {
		root, tail *node[T]

		len uint
	}
)

func newSO[T any](compare func(a, b T) bool) *singlyOrdered[T] {
	return &singlyOrdered[T]{}
}

func (l *singlyOrdered[T]) Push(body T, compare func(a, b T) bool) {
	node := &node[T]{
		body: body,
	}

	if l.root == nil {
		l.root = node
		l.tail = node

		l.len++

		return
	}

	last := l.root
	for n := l.root; n != nil; n = n.next {
		if compare(n.body, body) {
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

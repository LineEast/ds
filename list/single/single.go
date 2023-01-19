package single

type (
	Node[T any] struct {
		Next *Node[T]
		Data T
	}

	List[T any] struct {
		Head, Tail *Node[T]
	}
)

func (list *List[T]) PushTail(data T) {
	tail := list.Tail

	node := &Node[T]{
		Data: data,
	}

	if tail != nil {
		tail.Next = node
	} else {
		list.Head = node
	}

	list.Tail = node
}

func (list *List[T]) PopHead() (data T, ok bool) {
	head := list.Head

	if head == nil {
		return
	}

	list.Head = head.Next

	data = head.Data
	ok = true
	return
}

func (list *List[T]) Each(action func(data T)) {
	node := list.Head

	for node != nil {
		action(node.Data)

		node = node.Next
	}
}

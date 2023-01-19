package list

type (
	Node[T any] struct {
		Previous, Next *Node[T]
		Data           T
	}

	List[T any] struct {
		Head, Tail *Node[T]
	}
)

func (list *List[T]) PushHead(data T) {
	head := list.Head

	node := &Node[T]{
		Next: head,
		Data: data,
	}

	if head != nil {
		head.Previous = node
	} else {
		list.Tail = node
	}

	list.Head = node
}

func (list *List[T]) PushTail(data T) {
	tail := list.Tail

	node := &Node[T]{
		Previous: tail,
		Data:     data,
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

func (list *List[T]) PopTail() (data T, ok bool) {
	tail := list.Tail

	if tail == nil {
		return
	}

	list.Tail = tail.Previous

	data = tail.Data
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

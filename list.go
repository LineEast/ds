package data

type (
	ListNode[T any] struct {
		Previous, Next *ListNode[T]

		Value T
	}

	List[T any] struct {
		Head, Tail *ListNode[T]
	}
)

func NewList[T any]() (list *List[T]) {
	list = &List[T]{}

	return
}

func (list *List[T]) PushHead(object T) {
	head := list.Head

	node := &ListNode[T]{
		Next:  head,
		Value: object,
	}

	if head != nil {
		head.Previous = node
	} else {
		list.Tail = node
	}

	list.Head = node
}

func (list *List[T]) PushTail(object T) {
	tail := list.Tail

	node := &ListNode[T]{
		Previous: tail,
		Value:    object,
	}

	if tail != nil {
		tail.Next = node
	} else {
		list.Head = node
	}

	list.Tail = node
}

func (list *List[T]) Each(action func(object T)) {
	node := list.Head

	for node != nil {
		action(node.Value)

		node = node.Next
	}
}

func (list *List[T]) Pop() (object T, ok bool) {
	tail := list.Tail

	if tail == nil {
		return
	}

	list.Tail = tail.Previous

	object = tail.Value
	ok = true
	return
}

package double

type (
	List[T any] struct {
		Head, Tail *Node[T]
	}

	Node[T any] struct {
		Previous, Next *Node[T]
		Data           T
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

func (list *List[T]) PopHead() (data T, exists bool) {
	head := list.Head

	if head == nil {
		return
	}

	list.Head = head.Next

	data = head.Data
	exists = true
	return
}

func (list *List[T]) PopTail() (data T, exists bool) {
	tail := list.Tail

	if tail == nil {
		return
	}

	list.Tail = tail.Previous

	data = tail.Data
	exists = true
	return
}

func (list *List[T]) Each(action func(data T)) {
	node := list.Head

action:
	action(node.Data)
	node = node.Next

	if node != nil {
		goto action
	}
}

func (list *List[T]) Find(equal func(data, clue T) bool, clue T) (node *Node[T], found bool) {
	node = list.Head

equal:
	if node == nil {
		return
	}

	found = equal(node.Data, clue)

	if found == false {
		node = node.Next
		goto equal
	}

	return
}

func (node *Node[T]) PushHead(list *List[T]) {
	head := list.Head

	if head != nil {
		head.Previous = node
	} else {
		list.Tail = node
	}

	list.Head = node
}

func (node *Node[T]) PushTail(list *List[T]) {
	tail := list.Tail

	if tail != nil {
		tail.Next = node
	} else {
		list.Head = node
	}

	list.Tail = node
}

func (node *Node[T]) PopHead(list *List[T]) bool {
	node = list.Head

	if node == nil {
		return false
	}

	list.Head = node.Next
	node.Next = nil

	return true
}

func (node *Node[T]) PopTail(list *List[T]) bool {
	node = list.Tail

	if node == nil {
		return false
	}

	list.Tail = node.Previous
	node.Previous = nil

	return true
}

func (node *Node[T]) Remove(list *List[T]) {
	head := list.Head
	tail := list.Tail

	if head == nil {
		goto clear
	}

	if head == node {
		list.Head = head.Next

		goto clear
	}

check:
	if head.Next == nil {
		goto clear
	}

	head = head.Next

	if head == node {
		head.Previous.Next = head.Next

		if head == tail {
			list.Tail = head.Previous
		} else {
			head.Next.Previous = head.Previous
		}

		goto clear
	}

	goto check

clear:
	node.Previous = nil
	node.Next = nil
}

package doublylinkedlist

func (list *List[T]) InsertFirstAny(a T) {
	node := &Node[T]{Data: a}
	list.InsertFirst(node)
}

func (list *List[T]) InsertFirst(node *Node[T]) {
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
		list.head.next.prev = list.head
	}

	list.len++
}

func (list *List[T]) InsertLastAny(a T) {
	node := &Node[T]{Data: a}
	list.InsertLast(node)
}

func (list *List[T]) InsertLast(node *Node[T]) {
	if list.len < 1 {
		list.InsertFirst(node)
	} else {
		tmp := list.tail
		list.tail = &Node[T]{Data: node.Data, prev: tmp}
		tmp.next = list.tail
		list.len++
	}
}

package singlelinkedlist

func (list *List[T]) InsertFirst(node *Node[T]) {
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.len++
}

func (list *List[T]) InsertFirstAny(a T) {
	node := &Node[T]{Data: a}
	list.InsertFirst(node)
}

func (list *List[T]) InsertLast(node *Node[T]) {
	if list.head == nil {
		list.InsertFirst(node)
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	list.len++
}

func (list *List[T]) InsertLastAny(a T) {
	node := &Node[T]{Data: a}
	list.InsertLast(node)
}

func (list *List[T]) InsertIndexAny(a T, i uint) {
	node := &Node[T]{Data: a}
	list.InsertIndex(node, i)
}

func (list *List[T]) InsertIndex(node *Node[T], i uint){
	if i > list.len + 1 || i == 0 {
		DisplayError("Index must be in range of 0 and length + 1")
		return
	}
	if i == 1 {
		list.InsertFirst(node)
		return
	}
	currentPrev := list.NodeAt(i - 1)
	current := list.NodeAt(i)
	currentPrev.next = node
	currentPrev.next.next = current
	list.len++
}
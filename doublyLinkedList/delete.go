package doublylinkedlist

func (list *List[T]) NodeAt(i uint) *Node[T] {
	current := list.head
	for i > 1 {
		current = current.next
		i--
	}
	return current
}

func (list *List[T]) DeleteFirst() {
	if list.len == 0 {
		DisplayError("Empty List")
	} else {
		list.head = list.head.next
		list.len--
	}
}

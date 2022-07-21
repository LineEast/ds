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
		list.head.prev = nil
		list.len--
	}
}

func (list *List[T]) DeleteLast() {
	if list.len < 1 {
		DisplayError("Empty List")
	} else {
		list.tail.prev.next = nil
		list.len--
	}
}

func (list *List[T]) DeleteIndex(i uint) {
	if list.len < 1 || list.len < i {
		DisplayError("Empty List or Index out of range")
	} else if list.len > 1 && list.len == i {
		list.DeleteLast()
	} else if list.len == 1 || i == 1 {
		list.DeleteFirst()
	} else {
		tmp := list.NodeAt(i - 1)
		tmp.next = tmp.next.next
		list.len--
	}
}

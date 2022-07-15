package singlelinkedlist

func (list *List) NodeAt(i uint) *Node {
	current := list.head
	for i > 1 {
		current = current.next
		i--
	}
	return current
}

func (list *List) DeleteIndex(i uint) {
	if i > list.len || i == 0 {
		DisplayError("Index must be less than length and more than 0")
		return
	}
	if list.len == i {
		list.DeleteLast()
	} else if i == 1 {
		list.DeleteFirst()
	} else {
		currentPrev := list.NodeAt(i - 1)
		currentPrev.next = list.NodeAt(i).next
		list.len--
	}
}

func (list *List) DeleteFirst() {
	if list.head.next == nil {
		list.head = nil
	} else {
		list.head = list.head.next
	}
	list.len--
}

func (list *List) DeleteLast() {
	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil

	list.len--
}

package doublylinkedlist

// import "log"

func (list *List[T]) InsertFirst(node T) {
	n := &Node[T]{Data: node}

	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		n.next = list.head
		list.head = n
		list.head.next.prev = list.head
		// log.Println(list.tail)
	}

	list.len++
}

func (list *List[T]) InsertLast(node T) {
	if list.len < 1 {
		list.InsertFirst(node)
	} else {
		t := list.tail

		list.tail = &Node[T]{Data: node, prev: t}

		t.next = list.tail


		list.len++
	}
}

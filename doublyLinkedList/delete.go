package doublylinkedlist

func (list *List[T]) NodeAt(i uint) *Node[T] {
	current := list.head
	for i > 1 {
		current = current.next
		i--
	}
	return current
}
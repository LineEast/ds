package doublylinkedlist

func (list *List[T]) Replace(node *Node[T], i uint) {
	if list.len < 1 || list.len < i {
		DisplayError("Empty List or Index out of range")
	} else {
		list.NodeAt(i).Data = node.Data
	}
}

func (list *List[T]) ReplaceAny(a T, i uint) {
	node := &Node[T]{Data: a}
	list.Replace(node, i)
}
package singlelinkedlist

func (list *List[T]) ReplaceAny(a T, i uint) {
	node := &Node[T]{Data: a}
	list.Replace(node, i)
	
}

func (list *List[T]) Replace(node *Node[T], i uint) {
	if i > list.len || i == 0 {
		DisplayError("Index must be less than length and more than 0")
	} else {
		list.NodeAt(i).Data = node.Data
	}
}
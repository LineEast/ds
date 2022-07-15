package singlelinkedlist

func (list *List) ReplaceAny(a any, i uint) {
	node := &Node{Data: a}
	list.Replace(node, i)
	
}

func (list *List) Replace(node *Node, i uint) {
	if i > list.len || i == 0 {
		DisplayError("Index must be less than length and more than 0")
		return
	}
	list.NodeAt(i).Data = node.Data
}
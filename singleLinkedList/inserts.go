package singlelinkedlist

func (list *List) InsertFirst(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.len++
}

func (list *List) InsertFirstAny(a any) {
	node := &Node{Data: a}
	list.InsertFirst(node)
}

func (list *List) InsertLast(node *Node) {
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

func (list *List) InsertLastAny(a any) {
	node := &Node{Data: a}
	list.InsertLast(node)
}

func (list *List) InsertIndexAny(a any, i uint) {
	node := &Node{Data: a}
	list.InsertIndex(node, i)
}

func (list *List) InsertIndex(node *Node, i uint){
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
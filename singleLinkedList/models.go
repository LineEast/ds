package singlelinkedlist

type (
	Node struct {
		Data any
		next *Node
	}

	List struct {
		len  uint
		head *Node
	}
)

package singlelinkedlist

type (
	Node[T any] struct {
		Data T
		next *Node[T]
	}

	List[T any] struct {
		len  uint
		head *Node[T]
	}
)

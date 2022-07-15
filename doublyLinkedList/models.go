package doublylinkedlist

type (
	Node[T any] struct {
		Data       T
		prev, next *Node[T]
	}

	List[T any] struct {
		len        uint
		head, tail *Node[T]
	}
)

package general

type (
	Node[T any] struct {
		Previous, Next *Node[T]
		Data           T
	}

	Tree[T any] struct {
		Head, Tail *Node[T]
	}
)

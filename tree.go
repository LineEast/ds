package data

type (
	Node[T any] struct {
		This T

		Parent   *Node[T]
		Children []Node[T]
	}

	Tree[T any] struct {
		This T

		Children []Node[T]
	}
)

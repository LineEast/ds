package btree

type (
	Tree[T any] Node[T]

	Node[T any] struct {
		This T

		Parent   *Node[T]
		Children []Node[T]
	}
)

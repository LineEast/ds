package btree

type (
	Tree[t any] Node[t]

	Node[t any] struct {
		This t

		Parent   *Node[t]
		Children []Node[t]
	}
)

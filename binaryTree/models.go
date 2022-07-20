package binarytree

type (
	Node [T any] struct {
		Data T
		nodeL, nodeR *Node[T]
	}
	Tree [T any] struct {
		len uint
		root *Node[T]
	}
)
package binarytreeanybeta

type Tree[T any] struct {
	Data         T
	NodeL, NodeR *Tree[T]

	Compare func(one, two T) int
}

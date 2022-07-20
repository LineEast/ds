package binarytree

func Init[T any]() *Tree[T] {
	return new(Tree[T])
}

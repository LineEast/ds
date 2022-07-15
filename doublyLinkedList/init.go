package doublylinkedlist

func Init[T any]() *List[T]{
	return new(List[T])
}
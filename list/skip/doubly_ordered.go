package skip

type (
	doublyOrdered[T any] struct {
		head, tail *node[T]

		level uint
	}
)

func NewDO[T any]() *doublyOrdered[T] {
	return &doublyOrdered[T]{}
}

func (l *doublyOrdered[T]) Push(b T) (node *node[T]) {

	return
}

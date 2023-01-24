package skip

type (
	node[T any] struct {
		body T

		next, prev, down *node[T]
	}
)

func (n *node[T]) Body() T {
	return n.body
}

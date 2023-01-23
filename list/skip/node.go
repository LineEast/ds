package skip

type (
	node[T any] struct {
		body T

		next []*node[T]
	}
)

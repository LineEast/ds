package skip

type (
	node[T any] struct {
		body T

		next, prev, down *node[T]
	}
)

package data

type (
	Item[T any] struct {
		This T

		Previous *T
		Next     *T
	}

	List[T any] struct {
		Current *T

		Children []Item[T]
	}
)

package queue

type (
	Queue[T any] struct {
		items []T
	}
)

// New is a constructor for a new Queue with start size.
func New[T any](size uint) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, size),
	}
}

// Len returns len of q
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// Put item into q
func (q *Queue[T]) Put(item T) {
	q.items = append(q.items, item)
}

// Get element from queue
func (q *Queue[T]) Get() (i T) {
	if q.IsEmpty() {
		return
	}

	return q.items[0]
}

// Get element and remove from the queue
func (q *Queue[T]) Pop() (item T) {
	item = q.items[0]
	q.items = q.items[1:]
	return
}

// Clear cleans the q queue
func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
}

// Empty returns a bool indicating if this queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}

	return false
}

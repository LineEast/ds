package queue

import "github.com/LineEast/ds/list/singly"

type (
	QueueList[T any] struct {
		items singly.List[T]
	}
)

func New[T any]() *QueueList[T] {
	return &QueueList[T]{}
}

func (q *QueueList[T]) Len() uint {
	return q.items.Len()
}

func (q *QueueList[T]) Put(item T) {
	q.items.PushFront(item)
}

func (q *QueueList[T]) Get() (i T) {
	if q.items.IsEmpty() {
		return
	}

	return q.items.Back().Body()
}

func (q *QueueList[T]) Pop() T {
	return q.items.PopBack()
}

func (q *QueueList[T]) Clear() {
	q.items.Clear()
}

func (q *QueueList[T]) IsEmpty() bool {
	return q.items.IsEmpty()
}

package stack

import (
	"github.com/LineEast/ds/list/singly"
)

type (
	Stack_List[T any] struct {
		stack singly.List[T]
	}
)

func NewStackList[T any]() *Stack_List[T] {
	return &Stack_List[T]{}
}

func (s *Stack_List[T]) Len() uint {
	return s.stack.Len()
}

func (s *Stack_List[T]) IsEmpty() bool {
	return s.stack.IsEmpty()
}

func (s *Stack_List[T]) Clear() {
	s.stack.Clear()
}

func (s *Stack_List[T]) Push(value T) {
	s.stack.PushFront(value)
}

func (s *Stack_List[T]) Pop() T {
	return s.stack.PopFront()
}

func (s *Stack_List[T]) Peek() T {
	return s.stack.Front().Body()
}

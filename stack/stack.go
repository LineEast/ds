package stack

type (
	Stack[T any] struct {
		stack []T
	}
)

// Create a new stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func NewWithSize[T any](size int) *Stack[T] {
	return &Stack[T]{
		stack: make([]T, 0, size),
	}
}

// IsEmpty: check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Get len of s stack
func (s *Stack[T]) Len() uint {
	return uint(len(s.stack))
}

// Clear s stack
func (s *Stack[T]) Clear() {
	s.stack = s.stack[:0]
}

// Push a new value onto the stack
func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

// Remove and return top element of stack.
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var result T
		return result
	}

	element := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return element
}

// View the top item on the stack
func (s *Stack[T]) Peek() T {
	return s.stack[s.Len()-1]
}

package doubly

type (
	Node[T any] struct {
		body T

		next, prev *Node[T]
	}
)

// Create new node with b body
func NewNode[T any](b T) *Node[T] {
	return &Node[T]{
		body: b,
	}
}

// Create empty node
func NewEmptyNode[T any]() *Node[T] {
	return &Node[T]{}
}

// Set b body to the n Node
func (n *Node[T]) SetBody(b T) {
	n.body = b
}

// Get next Node from n Node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Get body of node
func (n *Node[T]) Body() T {
	return n.body
}

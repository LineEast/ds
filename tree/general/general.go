package general

import (
	"github.com/go-asphyxia/data/list/single"
)

type (
	Tree[T any] struct {
		Children *single.List[Node[T]]
	}

	Node[T any] struct {
		Children *single.List[Node[T]]
		Data     T
	}
)

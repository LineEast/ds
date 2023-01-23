package skip

import (
	"math/rand"
	"time"
)

type (
	SkipList[T any] struct {
		levels []*singlyOrdered[T]

		maxLevel uint
		rand     *rand.Rand

		compare func(a, b T) bool
	}
)

func New[T any](compare func(a, b T) bool, maxlvl uint) *SkipList[T] {
	source := rand.NewSource(time.Now().UnixNano())

	return &SkipList[T]{
		levels:   make([]*singlyOrdered[T], 0, maxlvl),
		maxLevel: maxlvl,
		rand:     rand.New(source),
		compare:  compare,
	}
}

func (s *SkipList[T]) Front() T {
	return s.levels[0].root.body
}

func (s *SkipList[T]) Back() T {
	return s.levels[0].tail.body
}

func (s *SkipList[T]) Len() uint {
	return s.levels[0].len
}

func (s *SkipList[T]) Push(v T) {
	s.levels[0].Push(v)
}

// func (s *SkipList[T]) Search(v T) {
// 	for i := s.maxLevel; i >= 0; i-- {
// 		// if s.levels[i].List
// 	}
// }

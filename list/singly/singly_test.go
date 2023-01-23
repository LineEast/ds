package singly

import (
	"testing"

	// "github.com/davecgh/go-spew/spew"
	"github.com/go-asphyxia/data/list/single"
)

const (
	success = "\t\u2713\t"
	failed  = "\t\u2717\t"
)

func newSingleListLineEast(b *testing.B) *List[int] {
	l := New[int]()
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	return l
}

func newSingleGohryt(b *testing.B) single.List[int] {
	l := single.List[int]{}
	for i := 0; i < b.N; i++ {
		l.PushTail(i)
	}

	return l
}

func BenchmarkGohryt(b *testing.B) {
	g := newSingleGohryt(b)

	n := g.Tail
	b.ResetTimer()

	n.Remove(&g)
}

func BenchmarkLineEast(b *testing.B) {
	l := newSingleListLineEast(b)
	n := l.tail

	b.ResetTimer()

	l.Remove(n)
}

// func TestClear(t *testing.T) {
// 	l := newIntList(5)

// 	l.Clear()

// 	if l.IsEmpty() {
// 		t.Logf("%s", success)
// 	} else {
// 		spew.Dump(l)
// 		t.Errorf("%s", failed)
// 	}
// }

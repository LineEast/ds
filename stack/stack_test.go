package stack

import (
	s "github.com/golang-collections/collections/stack"
	"testing"
)

type (
	Test struct {
		num  int64
		str  string
		fl   float64
		stru struct {
			num int64
			str string
			fl  float64
		}
	}
)

var t = Test{
	num: 999999,
	str: "biiig striiiiing",
	fl:  0.0000000001,
}

func BenchmarkPushSliceWichSize(b *testing.B) {
	s := NewWithSize[int](100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(382399)
	}
}

func BenchmarkPushSlice(b *testing.B) {
	s := New[int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(382399)
	}
}

func BenchmarkPushStandart(b *testing.B) {
	s := s.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(382399)
	}
}

func BenchmarkPushList(b *testing.B) {
	s := NewStackList[int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(382399)
	}
}

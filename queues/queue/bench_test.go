package queue

import (
	"testing"
)

func BenchmarkMyQueue(b *testing.B) {
	q := New(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := q.Get()
		q.Put(o)
	}
}

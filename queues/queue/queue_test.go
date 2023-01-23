package queue

import (
	"testing"
)

func QueueTest(t *testing.T) {
	item := "some"
	q := New(256)
	q.Put(item)
	ritem := q.Get()
	t.Errorf("%s", ritem)
}

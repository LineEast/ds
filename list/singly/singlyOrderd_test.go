package singly

// import (
// 	"github.com/davecgh/go-spew/spew"
// 	"testing"
// )

// func compare(a, b int) bool {
// 	return a > b
// }

// func TestOrderedList_Remove(t *testing.T) {
// 	l := NewOrderdList(compare)

// 	a := NewNode(6)
// 	b := NewNode(7)
// 	c := NewNode(8)

// 	l.PushNode(a)
// 	l.PushNode(b)
// 	l.PushNode(c)

// 	// ============================================================
// 	l.List.Remove(a)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove A len is not 2")
// 	}

// 	if l.List.Front().Body() != b.Body() {
// 		t.Errorf("%s\t%s", failed, "remove A front")
// 	}

// 	if l.List.Back().Body() != c.Body() {
// 		t.Errorf("%s\t%s", failed, "remove A back")
// 	}

// 	l.PushNode(a)
// 	// ============================================================
// 	l.List.Remove(b)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove B len is not 2")
// 	}

// 	if l.List.Front().Body() != a.Body() {
// 		t.Errorf("%s\t%s", failed, "remove B front")
// 	}

// 	if l.List.Back().Body() != c.Body() {
// 		t.Errorf("%s\t%s", failed, "remove B back")
// 	}

// 	if l.List.Front().Next().Body() != c.Body() {
// 		t.Errorf("%s\t%s", failed, "remove B error")
// 	}

// 	l.PushNode(b)
// 	// ============================================================
// 	l.List.Remove(c)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove C len is not 2")
// 	}

// 	if l.List.Front().Body() != a.Body() {
// 		t.Errorf("%s\t%s", failed, "remove C front")
// 	}

// 	if l.List.Back().Body() != b.Body() {
// 		t.Errorf("%s\t%s", failed, "remove C back")
// 	}

// 	l.PushNode(c)
// 	// ============================================================

// 	t.Log(success)
// }

// func compareEqual[T int](a, b T) bool {
// 	return a == b
// }

// func TestOrderedList_RemoveByBody(t *testing.T) {
// 	l := NewOrderdList(compare)
// 	l.Push(6)
// 	l.Push(7)
// 	l.Push(8)

// 	ce := compareEqual[int]
// 	// ============================================================
// 	l.List.RemoveByBody(ce, 6)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove A len is not 2")
// 	}

// 	if l.List.Front().Body() != 7 {
// 		t.Errorf("%s\t%s", failed, "remove A front")
// 	}

// 	if l.List.Back().Body() != 8 {
// 		t.Errorf("%s\t%s", failed, "remove A back")
// 	}

// 	l.Push(6)
// 	// ============================================================
// 	l.List.RemoveByBody(ce, 7)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove B len is not 2")
// 	}

// 	if l.List.Front().Body() != 6 {
// 		t.Errorf("%s\t%s", failed, "remove B front")
// 	}

// 	if l.List.Back().Body() != 8 {
// 		t.Errorf("%s\t%s", failed, "remove B back")
// 	}

// 	if l.List.Front().Next().Body() != 8 {
// 		t.Errorf("%s\t%s", failed, "remove B error")
// 	}

// 	l.Push(7)
// 	// ============================================================
// 	l.List.RemoveByBody(ce, 8)

// 	if l.List.Len() != 2 {
// 		t.Errorf("%s\t%s", failed, "remove C len is not 2")
// 	}

// 	if l.List.Front().Body() != 6 {
// 		t.Errorf("%s\t%s", failed, "remove C front")
// 	}

// 	if l.List.Back().Body() != 7 {
// 		t.Errorf("%s\t%s", failed, "remove C back")
// 	}

// 	l.Push(8)
// 	// ============================================================

// 	spew.Dump(l)
// 	t.Log(success)
// }

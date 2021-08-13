package solution

import "testing"

type Int int

func (i Int) Less(j QItem) bool {
	return i < j.(Int)
}
func Test_NewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Push(Int(5))
	pq.Push(Int(8))
	pq.Push(Int(3))

	first := pq.Front()
	if first != Int(3) {
		t.Error("first should be 3")
		return
	}

	first = pq.Pop()
	if first != Int(3) {
		t.Error("first should be 3")
		return
	}

	second := pq.Pop()
	if second != Int(5) {
		t.Error("second should be 5")
		return
	}

	pq.Push(Int(1))
	length := pq.Length()
	if length != 2 {
		t.Error("length should be 2")
		return
	}

	third := pq.Front()
	if third != Int(1) {
		t.Error("third should be 1")
		return
	}

	third = pq.Pop()
	if third != Int(1) {
		t.Error("third should be 1")
		return
	}

	fourth := pq.Pop()
	if fourth != Int(8) {
		t.Error("fourth should be 8")
		return
	}

	length = pq.Length()
	if length != 0 {
		t.Error("empty length should be 0")
		return
	}
}

package SeatManager

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	ih    *IntHeap
	count int
}

func Constructor(n int) SeatManager {
	ih := &IntHeap{}
	return SeatManager{ih: ih, count: 1}
}

func (sm *SeatManager) Reserve() int {
	if sm.ih.Len() == 0 {
		res := sm.count
		sm.count++
		return res
	} else {
		return heap.Pop(sm.ih).(int)
	}
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	heap.Push(sm.ih, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

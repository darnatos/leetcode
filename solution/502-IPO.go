package solution

import "container/heap"

type maxHeap [][]int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Pop() (x interface{}) {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	caph := make(maxHeap, 0, n)
	proh := make(maxHeap, 0, n)
	for i := 0; i < n; i++ {
		if capital[i] <= w {
			heap.Push(&proh, []int{profits[i]})
		} else {
			heap.Push(&caph, []int{-capital[i], profits[i]})
		}
	}

	for k > 0 && proh.Len() > 0 {
		cur := heap.Pop(&proh).([]int)
		w += cur[0]
		k--
		for caph.Len() > 0 && -caph[0][0] <= w {
			heap.Push(&proh, []int{caph[0][1]})
			heap.Pop(&caph)
		}
	}
	return w
}

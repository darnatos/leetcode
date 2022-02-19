package solution

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumDeviation(nums []int) int {
	// n:=len(nums)
	// even => x=x/2 => until x become an even number
	// odd => y=y*2 => 1 time
	// deviation => max(nums)-min(nums)
	// alg: max heap
	// TC: O(nlog(n)*log(k)), where n = len(nums) and k = max(nums)
	// SC: O(n) for heap
	hp := make(MaxHeap, 0)
	l := math.MaxInt32
	for _, j := range nums {
		if j&1 == 1 {
			j *= 2
		}
		heap.Push(&hp, j)
		l = min(l, j)
	}
	res := hp[0] - l
	for hp[0]&1 == 0 {
		top := hp[0] / 2
		heap.Pop(&hp)
		heap.Push(&hp, top)
		l = min(l, top)
		res = min(res, hp[0]-l)
	}
	return res
}

package solution

import (
	"container/heap"
	"github.com/darnatos/leetcode/util"
	"math"
)

type MinHeap []int

func (ih MinHeap) Len() int           { return len(ih) }
func (ih MinHeap) Less(i, j int) bool { return ih[i] < ih[j] }
func (ih MinHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }
func (ih *MinHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[:n-1]
	return x
}
func (ih *MinHeap) Peep() interface{} {
	o := *ih
	return o[0]
}

func (ih *MinHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	var cur int64 = 0
	suf := make([]int64, 3*n)
	pq := make(MinHeap, 0, n+1)
	for i := 3*n - 1; i >= 0; i-- {
		cur += int64(nums[i])
		heap.Push(&pq, nums[i])
		if pq.Len() > n {
			cur -= int64(pq[0])
			heap.Pop(&pq)
		}
		suf[i] = cur
	}

	pq = pq[:0]
	cur = 0
	var res int64 = math.MaxInt64

	for i := 0; i < 2*n; i++ {
		cur += int64(nums[i])
		heap.Push(&pq, -nums[i])
		if pq.Len() > n {
			cur -= -int64(pq[0])
			heap.Pop(&pq)
		}

		if pq.Len() == n {
			res = util.Min64(res, cur-suf[i+1])
		}
	}

	return res
}

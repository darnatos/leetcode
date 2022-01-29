package solution

import "container/heap"

func busiestServers(k int, arrival []int, load []int) []int {
	reqCount := make([]int, k)
	busy := make(PairPQ, 0)
	available := make(PairPQ, 0)
	for i := 0; i < k; i++ {
		heap.Push(&available, &Pair{i, 0})
	}

	for i := range arrival {
		cur := arrival[i]
		for busy.Len() > 0 && busy[0].first <= cur {
			x := heap.Pop(&busy).(*Pair).second
			heap.Push(&available, &Pair{i + (k+(x-i)%k)%k, 0})
		}
		if available.Len() > 0 {
			srv := heap.Pop(&available).(*Pair).first % k
			reqCount[srv]++
			heap.Push(&busy, &Pair{cur + load[i], srv})
		}
	}

	ma := 0
	for i := range reqCount {
		if ma < reqCount[i] {
			ma = reqCount[i]
		}
	}
	var res []int
	for i := range reqCount {
		if reqCount[i] == ma {
			res = append(res, i)
		}
	}
	return res
}

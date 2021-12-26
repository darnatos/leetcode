package solution

import "container/heap"

type PPQ [][]int

func (pq PPQ) Len() int {
	return len(pq)
}

func (pq PPQ) Less(i, j int) bool {
	return pq[i][0]*pq[i][0]+pq[i][1]*pq[i][1] > pq[j][0]*pq[j][0]+pq[j][1]*pq[j][1]
}

func (pq PPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PPQ) Push(x interface{}) {
	item := x.([]int)
	*pq = append(*pq, item)
}

func (pq *PPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	res := make(PPQ, 0, k)
	heap.Init(&res)

	for i := range points {
		heap.Push(&res, points[i])
		if res.Len() > k {
			heap.Pop(&res)
		}
	}

	return res
}

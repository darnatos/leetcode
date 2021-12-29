package solution

import "container/heap"

type KspPQ [][]int

func (pq KspPQ) Len() int {
	return len(pq)
}
func (pq KspPQ) Less(i, j int) bool {
	return pq[i][0]+pq[i][1] < pq[j][1]+pq[j][0]
}
func (pq KspPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *KspPQ) Push(x interface{}) {
	item := x.([]int)
	*pq = append(*pq, item)
}
func (pq *KspPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := make(KspPQ, 0, k)
	heap.Init(&pq)

	res := make([][]int, 0, k)
	for i := 0; i < min(len(nums1), k); i++ {
		heap.Push(&pq, []int{nums1[i], nums2[0], 0})
	}
	for len(pq) > 0 && len(res) < k {
		cur := heap.Pop(&pq).([]int)
		res = append(res, cur[:2])
		if cur[2] >= len(nums2)-1 {
			continue
		}
		heap.Push(&pq, []int{cur[0], nums2[cur[2]+1], cur[2] + 1})
	}

	return res
}

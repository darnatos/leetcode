package solution

import "container/heap"

type TripletPQ [][]int

func (pq TripletPQ) Len() int {
	return len(pq)
}
func (pq TripletPQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}
func (pq TripletPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *TripletPQ) Push(x interface{}) {
	item := x.([]int)
	*pq = append(*pq, item)
}
func (pq *TripletPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
func kSmallestTriplets(nums1 []int, nums2 []int, k int) []int {
	pq := make(TripletPQ, 0, k)
	heap.Init(&pq)

	res := make([]int, 0, k)
	for i := 0; i < min(len(nums1), k); i++ {
		heap.Push(&pq, []int{nums1[i] + nums2[0], nums1[i], nums2[0], 0})
	}
	for len(pq) > 0 && k > 0 {
		cur := heap.Pop(&pq).([]int)
		res = append(res, cur[0])
		if cur[3] >= len(nums2)-1 {
			continue
		}
		heap.Push(&pq, []int{cur[1] + nums2[cur[3]+1], cur[1], nums2[cur[3]+1], cur[3] + 1})
		k--
	}

	return res
}
func kthSmallest(mat [][]int, k int) int {
	res := mat[0]
	for i := 1; i < len(mat); i++ {
		res = kSmallestTriplets(res, mat[i], k)
	}
	return res[k-1]
}

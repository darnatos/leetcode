package SingleThreadCPU

import (
	"container/heap"
	"sort"
)

type Task struct {
	idx            int
	enqueueTime    int
	processingTime int
}
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a].processingTime < pq[b].processingTime || pq[a].processingTime == pq[b].processingTime && pq[a].idx < pq[b].idx
}
func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return x
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Task))
}

func GetOrder(tasks [][]int) []int {
	itasks := make([][]int, len(tasks))
	for i, t := range tasks {
		itasks[i] = []int{i, t[0], t[1]}
	}
	sort.Slice(itasks, func(i, j int) bool { return itasks[i][1] < itasks[j][1] })
	pq := make(PriorityQueue, 0, len(itasks))
	res := make([]int, 0, len(itasks))

	i, next := 0, itasks[0][0]
	for len(res) < len(itasks) {
		for i < len(itasks) && next >= itasks[i][1] {
			heap.Push(&pq, &Task{
				idx:            itasks[i][0],
				enqueueTime:    itasks[i][1],
				processingTime: itasks[i][2],
			})
			i++
		}
		if pq.Len() == 0 {
			next = itasks[i][1]
			continue
		}
		task := heap.Pop(&pq).(*Task)
		res = append(res, task.idx)
		next += task.processingTime
	}
	return res
}

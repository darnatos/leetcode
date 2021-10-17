package solution

import (
	"container/heap"
	"math"
)

type Pair struct {
	first  int
	second int
}
type PairPQ []*Pair

func (pq PairPQ) Len() int {
	return len(pq)
}
func (pq PairPQ) Less(i, j int) bool {
	return pq[i].second < pq[j].second
}
func (pq PairPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PairPQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}
func (pq *PairPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func SecondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	pq := &PairPQ{&Pair{1, 0}}
	heap.Init(pq)

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	visited := make([]int, n+1)
	visited[1] = 1

	for pq.Len() > 0 {
		top := heap.Pop(pq).(*Pair)
		city, timeElapsed := top.first, top.second

		if city == n && timeElapsed > dist[n] {
			return timeElapsed
		}
		if dist[city] == math.MaxInt32 {
			dist[city] = timeElapsed
		} else if dist[city] == timeElapsed || visited[city] >= 2 {
			continue
		} else {
			visited[city]++
		}

		// change  >  change  >  change  > ...
		// green   >   red    >  green   > ...
		// |--timeElapsed--|  |  <--- align to green light
		factor := timeElapsed / change
		if factor&1 == 1 {
			timeElapsed = (factor + 1) * change
		}

		nextTimeElapsed := timeElapsed + time
		for _, nextCity := range adj[city] {
			if visited[nextCity] < 2 {
				heap.Push(pq, &Pair{nextCity, nextTimeElapsed})
			}
		}
	}

	return -1
}

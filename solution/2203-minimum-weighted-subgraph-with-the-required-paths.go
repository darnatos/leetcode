package solution

import (
	"container/heap"
)

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	graph := make([][][]int, n)
	graph2 := make([][][]int, n)
	for _, e := range edges {
		graph[e[1]] = append(graph[e[1]], []int{e[0], e[2]})
		graph2[e[0]] = append(graph2[e[0]], []int{e[1], e[2]})
	}
	s1 := bfs(graph2, src1)
	s2 := bfs(graph2, src2)
	d := bfs(graph, dest)
	if d[src1] == 10000000000 || d[src2] == 10000000000 {
		return -1
	}
	var res int64 = 10000000000
	for i := 0; i < n; i++ {
		res = min64(res, d[i]+s1[i]+s2[i])
	}
	return res
}

func bfs(graph [][][]int, dest int) []int64 {
	q := make(pq, 0)
	heap.Push(&q, &pair{0, dest})
	dist := make([]int64, len(graph))
	for i := range dist {
		dist[i] = 10000000000
	}
	dist[dest] = 0
	for q.Len() > 0 {
		d, u := q[0].w, q[0].p
		heap.Pop(&q)
		if dist[u] != d {
			continue
		}
		for _, adj := range graph[u] {
			v, w := adj[0], int64(adj[1])
			if dist[v] > d+w {
				dist[v] = d + w
				heap.Push(&q, &pair{dist[v], v})
			}
		}
	}
	return dist
}

type pair struct {
	w int64
	p int
}
type pq []*pair

func (q pq) Len() int {
	return len(q)
}
func (q pq) Less(a, b int) bool {
	return q[a].w > q[b].w
}
func (q pq) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}
func (q *pq) Push(x interface{}) {
	*q = append(*q, x.(*pair))
}
func (q *pq) Pop() (x interface{}) {
	old := *q
	n := len(old)
	res := old[n-1]
	old = old[:n-1]
	*q = old
	return res
}

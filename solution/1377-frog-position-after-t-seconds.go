package solution

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	adj := make([][]int, n+1)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	q := make([]int, 0, n)
	visited := make([]bool, n+1)
	prob := make([]float64, n+1)
	var cur int
	var cnt float64

	q = append(q, 1)
	visited[1] = true
	prob[1] = 1.0
	for len(q) > 0 && t > 0 {
		t--
		end := len(q)
		for i := 0; i < end; i++ {
			cur = q[i]
			cnt = 0
			for _, u := range adj[cur] {
				if !visited[u] {
					cnt++
				}
			}
			for _, u := range adj[cur] {
				if !visited[u] {
					visited[u] = true
					prob[u] = prob[cur] / cnt
					q = append(q, u)
				}
			}
			if cnt > 0 {
				prob[cur] = 0
			}
		}
		q = q[end:]
	}
	return prob[target]
}

package solution

func FindMinHeightTrees(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	//find highest rank node
	end := 0
	rank := make([]int, n)
	queue := []int{0}
	rank[0] = 1
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if rank[v] > 0 {
				continue
			}
			rank[v] = rank[u] + 1
			if rank[end] < rank[v] {
				end = v
			}
			queue = append(queue, v)
		}
	}

	//construct longest path
	rank2 := make([]int, n)
	rank2[end] = 1
	queue = []int{end}
	pre := make([]int, n)
	pre[end] = -1
	start := 0
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if rank2[v] > 0 {
				continue
			}
			rank2[v] = rank2[u] + 1
			if rank2[start] < rank2[v] {
				start = v
			}
			queue = append(queue, v)
			pre[v] = u
		}
	}

	path := make([]int, 0)
	for start != -1 {
		path = append(path, start)
		start = pre[start]
	}

	return path[(len(path)-1)/2 : len(path)/2+1]
}

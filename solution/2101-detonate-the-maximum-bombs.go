package solution

func maximumDetonation(bombs [][]int) int {
	adj := make([][]int, len(bombs))
	for i := range bombs {
		for j := range bombs {
			if i == j {
				continue
			}
			dx, dy := bombs[i][0]-bombs[j][0], bombs[i][1]-bombs[j][1]
			if bombs[i][2]*bombs[i][2] >= dx*dx+dy*dy {
				adj[i] = append(adj[i], j)
			}
		}
	}

	res := 0
	for i := 0; i < len(bombs); i++ {
		visited := make([]bool, len(bombs))
		res = max(res, mdDfs(adj, i, visited))
	}

	return res
}

func mdDfs(adj [][]int, u int, visited []bool) int {
	if visited[u] {
		return 0
	}
	visited[u] = true
	res := 1
	for _, v := range adj[u] {
		res += mdDfs(adj, v, visited)
	}
	return res
}

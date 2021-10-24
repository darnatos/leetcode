package solution

func CountHighestScoreNodes(parents []int) int {
	n := len(parents)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[parents[i]] = append(adj[parents[i]], i)
	}

	counts := make([]int, n)

	best, res := -1, 0
	var dfs func(int)
	dfs = func(node int) {
		counts[node] = 1
		cur := 1
		for _, i := range adj[node] {
			dfs(i)
			counts[node] += counts[i]
			cur *= counts[i]
		}
		if node > 0 {
			cur *= n - counts[node]
		}
		if best < cur {
			best = cur
			res = 1
		} else if cur == best {
			res++
		}

	}
	dfs(0)

	return res
}

package solution

func LoudAndRich(richer [][]int, quiet []int) []int {
	adj := make([][]int, len(quiet))
	for _, v := range richer {
		adj[v[1]] = append(adj[v[1]], v[0])
	}
	ans := make([]int, len(quiet))
	for i := range ans {
		ans[i] = -1
	}

	for i := range ans {
		dfsLoudAndRich(adj, quiet, i, ans)
	}

	return ans
}

func dfsLoudAndRich(adj [][]int, quiet []int, i int, ans []int) int {
	if ans[i] >= 0 {
		return ans[i]
	}
	ans[i] = i
	for _, j := range adj[i] {
		if quiet[ans[i]] > quiet[dfsLoudAndRich(adj, quiet, j, ans)] {
			ans[i] = ans[j]
		}
	}
	return ans[i]
}

func LoudAndRich2(richer [][]int, quiet []int) []int {
	N := len(quiet)
	adj := make([][]int, N)
	indegrees := make([]int, N)
	for _, v := range richer {
		adj[v[0]] = append(adj[v[0]], v[1])
		indegrees[v[1]]++
	}

	q := make([]int, 0)
	ans := make([]int, N)
	for i := range indegrees {
		if indegrees[i] == 0 {
			q = append(q, i)
		}
		ans[i] = i
	}

	for len(q) != 0 {
		f := q[0]
		q = q[1:]

		for _, i := range adj[f] {
			if quiet[i] > quiet[ans[f]] {
				ans[i] = ans[f]
				quiet[i] = quiet[ans[f]]
			}
			indegrees[i]--
			if indegrees[i] == 0 {
				q = append(q, i)
			}
		}
	}
	return ans
}
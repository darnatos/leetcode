package solution

func MaximalPathQuality(values []int, edges [][]int, maxTime int) int {
	graph := make([][][]int, len(values))
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], []int{e[1], e[2]})
		graph[e[1]] = append(graph[e[1]], []int{e[0], e[2]})
	}

	visited := make([]int, len(values))
	return mpqDfs(graph, visited, values, 0, maxTime, 0, 0)
}

func mpqDfs(graph [][][]int, visited, values []int, u, time, gain, res int) int {
	if visited[u] == 0 {
		gain += values[u]
	}
	visited[u]++
	if u == 0 {
		if res < gain {
			res = gain
		}
	}
	for _, neib := range graph[u] {
		if neib[1] <= time {
			res = mpqDfs(graph, visited, values, neib[0], time-neib[1], gain, res)
		}
	}
	visited[u]--
	return res
}

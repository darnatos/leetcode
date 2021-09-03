package solution

func CriticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)

	for i := range graph {
		graph[i] = make([]int, 0, 2)
	}
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	ranks := make([]int, n)
	for i := range ranks {
		ranks[i] = -2
	}

	res := make([][]int, 0, 4)
	ccDfs(graph, n, 0, 0, ranks, &res)

	return res
}

func ccDfs(graph [][]int, n, vertex, r int, ranks []int, res *[][]int) int {
	if ranks[vertex] != -2 {
		return ranks[vertex]
	}

	lR := r
	ranks[vertex] = r

	for _, v := range graph[vertex] {
		if ranks[v] == r-1 || ranks[v] == n {
			continue
		}

		vR := ccDfs(graph, n, v, r+1, ranks, res)
		lR = min(lR, vR)
		if vR > r {
			*res = append(*res, []int{vertex, v})
		}
	}

	ranks[vertex] = n
	return lR
}

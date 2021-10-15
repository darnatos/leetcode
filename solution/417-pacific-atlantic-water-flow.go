package solution

func PacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	graph := make([][]bool, m)
	graph2 := make([][]bool, m)
	queue := make([][]int, 0, m+n)
	queue2 := make([][]int, 0, m+n)
	for i := range graph {
		graph[i] = make([]bool, n)
		graph2[i] = make([]bool, n)
	}
	for j := range graph[0] {
		graph[0][j] = true
		graph2[m-1][j] = true
		queue = append(queue, []int{0, j})
		queue2 = append(queue2, []int{m - 1, j})
	}
	for i := range graph {
		graph[i][0] = true
		graph2[i][n-1] = true
		queue = append(queue, []int{i, 0})
		queue2 = append(queue2, []int{i, n - 1})
	}
	dir := [...]int{0, -1, 0, 1, 0}
	bfs := func(queue [][]int, graph [][]bool) {
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				x := cur[0] + dir[i]
				y := cur[1] + dir[i+1]
				if x < 0 || y < 0 || x >= m || y >= n || graph[x][y] || heights[x][y] < heights[cur[0]][cur[1]] {
					continue
				}
				queue = append(queue, []int{x, y})
				graph[x][y] = true
			}
		}
	}
	bfs(queue, graph)
	bfs(queue2, graph2)

	res := make([][]int, 0)
	for i := range heights {
		for j := range heights[0] {
			if graph[i][j] && graph2[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

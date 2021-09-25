package solution

func ShortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k >= m+n-2 {
		return m + n - 2
	}

	tmp := make([]int, m*n)
	graph := make([][]int, m)
	for i := range graph {
		graph[i] = tmp[n*i : n*(i+1)]
		for j := range graph[i] {
			graph[i][j] = -1
		}
	}

	d := [...]int{0, 1, 0, -1, 0}
	q := [][]int{{0, 0, 0, k}}
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		w, x, y, z := p[0], p[1], p[2], p[3]

		if x < 0 || x >= m || y < 0 || y >= n {
			continue
		}

		if x == m-1 && y == n-1 {
			return w
		}

		z -= grid[x][y]
		if z < 0 {
			continue
		}

		if graph[x][y] != -1 && graph[x][y] >= z {
			continue
		}
		graph[x][y] = z

		for i := 0; i < 4; i++ {
			q = append(q, []int{w + 1, x + d[i], y + d[i+1], z})
		}
	}
	return -1
}

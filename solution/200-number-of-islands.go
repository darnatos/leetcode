package solution

func NumIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	num := 0
	queue := make([][]int, 0, 16)

	AdjVector := [...]int{0, -1, 0, 1, 0}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			queue = append(queue, []int{i, j})
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]

				for k := 0; k < 4; k++ {
					x := node[0] + AdjVector[k]
					y := node[1] + AdjVector[k+1]
					if x < 0 || x >= m || y < 0 || y >= n {
						continue
					}
					if grid[x][y] == '1' {
						queue = append(queue, []int{x, y})
						grid[x][y] = '0'
					}
				}
			}
			num++
		}
	}

	return num
}

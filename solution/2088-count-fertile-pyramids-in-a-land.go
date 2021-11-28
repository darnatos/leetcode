package solution

func FindAllPeopleCountPyramids(grid [][]int) int {
	res := countPyramid(grid)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}
	res += countPyramid(grid)
	return res
}

func countPyramid(grid [][]int) int {
	n, m, res := len(grid), len(grid[0]), 0
	for i := 1; i < n; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] > 0 && grid[i-1][j] > 0 {
				grid[i][j] = min(grid[i-1][j-1], grid[i-1][j+1]) + 1
				res += grid[i][j] - 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] > 0 {
				grid[i][j] = 1
			}
		}
	}
	return res
}

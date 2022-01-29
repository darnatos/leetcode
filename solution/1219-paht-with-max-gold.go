package solution

func getMaximumGold(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[0] {
			res = max(res, gmgDfs(grid, i, j))
		}
	}
	return res
}

func gmgDfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	tmp := grid[i][j]
	grid[i][j] = 0
	m := max(max(gmgDfs(grid, i-1, j), gmgDfs(grid, i+1, j)), max(gmgDfs(grid, i, j-1), gmgDfs(grid, i, j+1)))
	grid[i][j] = tmp
	return tmp + m
}

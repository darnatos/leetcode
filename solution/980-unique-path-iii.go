package solution

func UniquePathsIII(grid [][]int) int {
	//grid[i][j]==2
	zeros := 1
	res := 0
	x, y := 0, 0

	for i := range grid {
		for j, e := range grid[i] {
			if e == 0 {
				zeros++
			} else if e == 1 {
				x, y = i, j
			}
		}
	}
	up3Dfs(grid, x, y, zeros, &res)

	return res
}

func up3Dfs(grid [][]int, x, y, zeros int, res *int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == -1 {
		return
	}
	if grid[x][y] == 2 {
		if zeros == 0 {
			*res++
		}
		return
	}
	grid[x][y] = -1

	up3Dfs(grid, x-1, y, zeros-1, res)
	up3Dfs(grid, x+1, y, zeros-1, res)
	up3Dfs(grid, x, y+1, zeros-1, res)
	up3Dfs(grid, x, y-1, zeros-1, res)

	grid[x][y] = 0
}

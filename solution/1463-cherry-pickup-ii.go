package solution

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	return cpDfs(grid, m, n, 0, 0, n-1, dp)
}

func cpDfs(grid [][]int, m, n, i, c1, c2 int, dp [][][]int) int {
	if i == m {
		return 0
	}
	if dp[i][c1][c2] != -1 {
		return dp[i][c1][c2]
	}
	res := 0
	for d1 := -1; d1 <= 1; d1++ {
		nc1 := c1 + d1
		if nc1 < 0 || nc1 >= n {
			continue
		}
		for d2 := -1; d2 <= 1; d2++ {
			nc2 := c2 + d2
			if nc2 < 0 || nc2 >= n {
				continue
			}
			res = max(res, cpDfs(grid, m, n, i+1, nc1, nc2, dp))
		}
	}
	cherries := grid[i][c1]
	if c1 != c2 {
		cherries += grid[i][c2]
	}
	dp[i][c1][c2] = res + cherries
	return dp[i][c1][c2]
}

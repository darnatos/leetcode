package solution

func UniquePathsWithObstacles(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = 1
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

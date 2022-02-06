package solution

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n+1)
	dp0 := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				dp[j] = dp[j-1] + grid[i-1][j-1]
				continue
			}
			if i > 1 && j == 1 {
				dp[j] = dp0[j] + grid[i-1][j-1]
				continue
			}
			dp[j] = min(dp[j-1], dp0[j]) + grid[i-1][j-1]
		}
		copy(dp0, dp)
	}
	return dp[n]
}

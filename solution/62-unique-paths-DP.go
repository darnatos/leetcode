package solution

func UniquePathsDP(m int, n int) int {
	dp0 := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		dp0, dp = dp, dp0
		dp[0] = 1
		for j := 1; j < n; j++ {
			dp[j] = dp0[j]+dp[j-1]
		}
	}
	return dp[n-1]
}
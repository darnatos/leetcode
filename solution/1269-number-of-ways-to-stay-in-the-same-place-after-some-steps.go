package solution

func numWays(steps int, arrLen int) int {
	ma := min((steps/2)+1, arrLen)
	dp, dp0 := make([]int, ma+2), make([]int, ma+2)
	dp[1] = 1
	for i := 1; i <= steps; i++ {
		dp0, dp = dp, dp0
		for j := 1; j <= min(i+1, ma); j++ {
			dp[j] = (dp0[j-1] + dp0[j] + dp0[j+1]) % 1000000007
		}
	}
	return dp[1]
}

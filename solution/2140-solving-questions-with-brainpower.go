package solution

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := range questions {
		if i > 0 {
			dp[i] = max64(dp[i], dp[i-1])
		}
		ni := min(n, i+questions[i][1]+1)
		dp[ni] = max64(dp[ni], int64(questions[i][0])+dp[i])
	}
	return dp[n]
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

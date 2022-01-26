package solution

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for k := range strs {
		zc, oc := 0, 0
		for _, ch := range strs[k] {
			if ch == '0' {
				zc++
			} else {
				oc++
			}
		}
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				if i+zc <= m && j+oc <= n {
					dp[i][j] = max(dp[i][j], dp[i+zc][j+oc]+1)
				}
			}
		}
	}
	return dp[0][0]
}

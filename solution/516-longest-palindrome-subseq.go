package solution

func longestPalindromeSubseq(s string) int {
	// dp[i][j]=dp[i+1][j-1]+2 if s[i]==s[j]
	n := len(s)
	dp := make([]int, n)
	dp0 := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp, dp0 = dp0, dp
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[j] = dp0[j-1] + 2
			} else {
				dp[j] = max(dp0[j], dp[j-1])
			}
		}
	}

	return dp[n-1]
}

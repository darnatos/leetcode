package solution

func NumDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	dp0 := make([]int, len(s)+1)
	dp := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(t); i++ {
		dp, dp0 = dp0, dp
		dp[0] = 0
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				dp[j+1] = dp0[j] + dp[j]
			} else {
				dp[j+1] = dp[j]
			}
		}
	}

	return dp[len(s)]
}

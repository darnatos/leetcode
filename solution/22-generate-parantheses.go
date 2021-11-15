package solution

func GenerateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, ls := range dp[j] {
				for _, rs := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+ls+")"+rs)
				}
			}
		}
	}
	return dp[n]
}

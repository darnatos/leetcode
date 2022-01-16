package solution

func numWays2(words []string, target string) int {
	m, n, k := len(words), len(words[0]), len(target)
	if n < k {
		return 0
	}
	cnt := make([][26]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt[j][words[i][j]-'a']++
		}
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		for j := 0; j <= i && j < k; j++ {
			dp[i+1][j+1] = ((dp[i][j]*cnt[i][target[j]-'a'])%1000000007 + dp[i][j+1]) % 1000000007
		}
	}
	return dp[n][k]
}

package solution

import "leetcode/util"

func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)

	arr := make([]int, (m+1)*(n+1))
	for i := 0; i <= m; i++ {
		dp[i] = arr[i*(n+1) : (i+1)*(n+1)]
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i-1][j], util.Min(dp[i-1][j-1], dp[i][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

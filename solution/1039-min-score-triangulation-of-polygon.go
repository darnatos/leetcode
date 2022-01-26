package solution

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+values[i]*values[k]*values[j]+dp[k][j])
			}
		}
	}
	return dp[0][n-1]
}

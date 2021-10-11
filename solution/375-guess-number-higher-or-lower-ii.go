package solution

import "math"

func GetMoneyAmount(n int) int {
	tmp := make([]int, (n+1)*(n+1))
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = tmp[i*(n+1) : (i+1)*(n+1)]
	}
	return gmaHelper(dp, 1, n)
}

func gmaHelper(dp [][]int, s, e int) int {
	if s >= e {
		return 0
	}
	if dp[s][e] != 0 {
		return dp[s][e]
	}
	cost := math.MaxInt32
	for i := s; i <= e; i++ {
		cost = min(cost, i+max(gmaHelper(dp, s, i-1), gmaHelper(dp, i+1, e)))
	}
	dp[s][e] = cost
	return cost
}
package solution

import (
	"math"
	"sort"
)

func minDistance(houses []int, K int) int {
	sort.Ints(houses)
	cost := make([][]int, len(houses))
	for i := range houses {
		cost[i] = make([]int, len(houses))
		for j := range houses {
			mPos := houses[i+(j-i)/2]
			for k := i; k <= j; k++ {
				cost[i][j] += abs(mPos - houses[k])
			}
		}
	}
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, len(houses))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dfsMd(cost, houses, len(houses), K, 0, dp)
}

func dfsMd(cost [][]int, houses []int, n, k, i int, dp [][]int) int {
	if k == 0 && i == n {
		return 0
	}
	if k == 0 || i == n {
		return math.MaxInt32
	}
	if dp[k][i] != -1 {
		return dp[k][i]
	}

	res := math.MaxInt32
	for j := i; j < n; j++ {
		res = min(res, cost[i][j]+dfsMd(cost, houses, n, k-1, j+1, dp))
	}
	dp[k][i] = res
	return res
}

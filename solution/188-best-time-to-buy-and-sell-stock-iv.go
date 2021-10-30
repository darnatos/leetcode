package solution

import (
	"github.com/darnatos/leetcode/util"
	"math"
)

func MaxProfit4(k int, prices []int) int {
	dp := make([]int, k+1)
	m := make([]int, k+1)
	for i := range m {
		m[i] = math.MaxInt32
	}

	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			m[j] = util.Min(m[j], prices[i]-dp[j-1])
			dp[j] = util.Max(dp[j], prices[i]-m[j])
		}
	}
	return dp[k]
}

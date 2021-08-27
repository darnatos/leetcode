package solution

import "leetcode/util"

func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	copy(dp, triangle[n-1])

	tmp := make([]int, n)

	for i := n-2; i >= 0; i-- {
		copy(tmp, dp)
		for j := range triangle[i] {
			dp[j] = triangle[i][j] + util.Min(tmp[j], tmp[j+1])
		}
	}
	return dp[0]
}
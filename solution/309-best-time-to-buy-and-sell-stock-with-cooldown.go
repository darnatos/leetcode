package solution

import "leetcode/util"

func MaxProfit(prices []int) int {
	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))
	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = 0

	for i := 1; i < len(prices); i++ {
		s0[i] = util.Max(s0[i-1], s2[i-1])
		s1[i] = util.Max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return util.Max(s0[len(prices)-1], s2[len(prices)-1])
}

package solution

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum < target {
		return 0
	}
	if (target+sum)%2 > 0 {
		return 0
	}

	s := (target + sum) >> 1
	dp := make(map[int]int, s+1)
	dp[0] = 1
	for _, n := range nums {
		for i := s; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[s]
}

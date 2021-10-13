package solution

func LastStoneWeightII(stones []int) int {
	dp := [1501]bool{}
	dp[0] = true
	sumA := 0
	for _, a := range stones {
		sumA += a
		for i := min(sumA, 1500); i >= a; i-- {
			dp[i] = dp[i] || dp[i-a]
		}
	}
	for i := sumA/2; i >= 0; i-- {
		if dp[i] {
			return sumA - i - i
		}
	}
	return 0
}

package solution

func MaxProfit3(prices []int) int {
	dp := [3]int{}
	m := [3]int{prices[0],prices[0],prices[0]}

	for i := 1; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			m[k] = min(m[k], prices[i]-dp[k-1])
			dp[k] = max(dp[k], prices[i]-m[k])
		}
	}
	return dp[2]
}

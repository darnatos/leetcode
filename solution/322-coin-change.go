package solution

func CoinChange(coins []int, amount int) int {
	m := amount + 1
	dp := make([]int, m)
	for i := range dp {
		dp[i] = m
	}
	dp[0] = 0

	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

package solution

func MaxProfit2(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0)
	}
	return res
}

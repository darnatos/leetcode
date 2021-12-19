package solution

func getDescentPeriods(prices []int) int64 {
	var pre int64 = 1
	var res int64 = 1
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			pre = pre + 1
		} else {
			pre = 1
		}
		res += pre
	}
	return res
}

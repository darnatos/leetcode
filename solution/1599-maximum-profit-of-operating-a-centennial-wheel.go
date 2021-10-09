package solution

func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	rest, ret, cur, max, i := 0, 0, 0, 0, 0
	for i < len(customers) || rest > 0 {
		if i < len(customers) {
			rest += customers[i]
		}
		i++
		cur -= runningCost
		if rest >= 4 {
			cur += 4 * boardingCost
			rest -= 4
		} else {
			cur += rest * boardingCost
			rest = 0
		}
		if max < cur {
			max = cur
			ret = i
		}
	}

	if max <= 0 {
		return -1
	}
	return ret
}

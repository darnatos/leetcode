package solution

func FindMinMoves(machines []int) int {
	total := 0
	for _, v := range machines {
		total += v
	}
	if total%len(machines) > 0 {
		return -1
	}

	avg, ret, pref := total/len(machines), 0, 0
	for _, v := range machines {
		ret = max(max(ret, abs(pref)), v-avg)
		pref += v - avg
	}
	return ret
}

package solution

func goodDaysToRobBank(security []int, time int) []int {
	pre := make([]int, len(security))
	for i := 0; i < len(security); i++ {
		if i > 0 && security[i] <= security[i-1] {
			pre[i] = 1 + pre[i-1]
		} else {
			pre[i] = 0
		}
	}
	suf := 1
	res := make([]int, 0, len(security))
	for i := len(security) - 1; i >= 0; i-- {
		if i < len(security)-1 && security[i] <= security[i+1] {
			suf++
		} else {
			suf = 0
		}
		if pre[i] >= time && suf >= time {
			res = append(res, i)
		}
	}

	return res
}

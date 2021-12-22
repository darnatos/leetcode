package solution

func numberOfWeeks(milestones []int) int64 {
	var sum int64 = 0
	m := 0
	for _, i := range milestones {
		sum += int64(i)
		if m < i {
			m = i
		}
	}
	return min64(sum, 2*(sum-int64(m))+1)
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

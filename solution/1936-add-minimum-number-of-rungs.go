package solution

func AddRungs(rungs []int, dist int) int {
	cur := 0
	count := 0
	for _, v := range rungs {
		count += (v - cur - 1) / dist
		cur = v
	}
	return count
}

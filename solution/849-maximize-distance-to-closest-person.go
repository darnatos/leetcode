package solution

func maxDistToClosest(seats []int) int {
	n := len(seats)
	pre := -1
	res := 0
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if pre < 0 {
				res = max(res, i)
			} else {
				res = max(res, (i-pre)/2)
			}
			pre = i
		}
	}
	res = max(n-1-max(0, pre), res)
	return res
}

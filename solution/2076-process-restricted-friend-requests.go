package solution

func FriendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	fr := make([]int, n)
	for i := range fr {
		fr[i] = i
	}

	res := make([]bool, len(requests))
	for i, req := range requests {
		u := find(fr, req[0])
		v := find(fr, req[1])
		if u > v {
			u, v = v, u
		}
		flag := true
		for _, res := range restrictions {
			x, y := find(fr, res[0]), find(fr, res[1])
			if x > y {
				x, y = y, x
			}
			if x == u && y == v {
				flag = false
				break
			}
		}
		if flag {
			union(fr, req[0], req[1])
		}
		res[i] = flag
	}
	return res
}

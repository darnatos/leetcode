package solution

type P struct {
	first  int
	second bool
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	f := make(map[int]bool)
	for i := range forbidden {
		f[forbidden[i]] = true
	}

	q := make([]P, 0, 64)
	q = append(q, P{0, false})
	res := 0
	for len(q) != 0 {
		s := len(q)
		for s > 0 {
			curr := q[0]
			q = q[1:]
			s--
			num := curr.first
			if num == x {
				return res
			}
			if f[num] == true {
				continue
			}
			f[num] = true
			if curr.second == false && num-b >= 0 {
				q = append(q, P{num - b, true})
			}
			if num <= 2000+b {
				q = append(q, P{num + a, false})
			}
		}
		res++
	}

	return -1
}

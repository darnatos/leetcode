package solution

func ShipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, w := range weights {
		if l < w {
			l = w
		}
		r += w
	}

	for l < r {
		m := l + (r-l)/2

		d, cur := 1, 0
		for _, w := range weights {
			cur += w
			if cur > m {
				d++
				cur = w
			}
		}

		if d > days {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

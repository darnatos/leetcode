package solution

func NthUglyNumber(n int, a int, b int, c int) int {
	ab := lcm(a, b)
	bc := lcm(b, c)
	ca := lcm(c, a)
	abc := lcm(ab, c)

	f := func(k int) int {
		return k/a + k/b + k/c - k/ab - k/bc - k/ca + k/abc
	}
	l, r := 1, int(2e9)
	for l < r {
		m := l + (r-l)/2
		num := f(m)
		if num < n {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

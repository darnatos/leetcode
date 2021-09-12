package solution

func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	prev, res := 1, 2
	for i := 2; i < n; i++ {
		prev, res = res, prev+res
	}
	return res
}

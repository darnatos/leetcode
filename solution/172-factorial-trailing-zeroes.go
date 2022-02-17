package solution

func trailingZeroes(n int) int {
	//number of 5^1, 5^2, ... , 5^k
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}

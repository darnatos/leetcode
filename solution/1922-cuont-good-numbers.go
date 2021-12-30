package solution

func countGoodNumbers(n int64) int {
	return int((modExp(n/2, 20) * (1 + 4*(n%2))) % mod)
}

func modExp(n, m int64) int64 {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return (modExp(n/2, (m*m)%mod)) % mod
	}
	return (m * modExp((n-1)/2, (m*m)%mod)) % mod
}

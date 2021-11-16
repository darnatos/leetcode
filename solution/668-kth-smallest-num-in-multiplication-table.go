package solution

func FindKthNumber(m int, n int, k int) int {
	lo, hi := 1, m*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		cnt := 0
		for i := 1; i <= m; i++ {
			cnt += min(mid/i, n)
		}
		if cnt >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

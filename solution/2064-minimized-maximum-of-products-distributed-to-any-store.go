package solution

func MinimizedMaximum(n int, quantities []int) int {
	lo, hi := 1, int(1e5)

	for lo < hi {
		mid := lo + (hi-lo)/2
		tot := 0
		for _, q := range quantities {
			tot += (q-1)/mid + 1
		}
		if tot > n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

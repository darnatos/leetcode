package solution

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	mod := 1
	for i := 1; i <= k; i++ {
		if mod%k == 0 {
			return i
		}
		mod = (mod*10 + 1) % k
	}
	return -1
}

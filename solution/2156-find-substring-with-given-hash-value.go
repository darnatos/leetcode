package solution

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)
	hv, pk := 0, 1
	start := 0
	val := func(i int) int {
		return int(s[i] - '`')
	}
	for i := n - 1; i >= 0; i-- {
		hv = (hv*power + val(i)) % modulo
		if i < n-k {
			hv = ((hv-pk*val(i+k))%modulo + modulo) % modulo
		} else {
			pk = pk * power % modulo
		}
		if hv == hashValue {
			start = i
		}
	}
	return s[start : start+k]
}

package solution

func LongestDiverseString(a int, b int, c int) string {
	return string(ldsHelper(a, b, c, 'a', 'b', 'c'))
}

func ldsHelper(a, b, c int, aa, bb, cc byte) []byte {
	if a < b {
		return ldsHelper(b, a, c, bb, aa, cc)
	}
	if b < c {
		return ldsHelper(a, c, b, aa, cc, bb)
	}
	if b == 0 {
		res := make([]byte, min(a, 2))
		for i := len(res) - 1; i >= 0; i-- {
			res[i] = aa
		}
		return res
	}
	ad := min(a, 2)
	bd := 1
	if b > a-ad {
		bd = 0
	}
	res := make([]byte, 0, a+b+c)
	for i := 0; i < ad; i++ {
		res = append(res, aa)
	}
	for i := 0; i < bd; i++ {
		res = append(res, bb)
	}
	res = append(res, ldsHelper(a-ad, b-bd, c, aa, bb, cc)...)
	return res
}

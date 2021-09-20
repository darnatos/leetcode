package solution

func NumSplits(s string) int {
	cnt, ls, rs := 0, make(map[byte]int), make(map[byte]int)

	for i := range s {
		rs[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if len(ls) == len(rs) {
			cnt++
		}

		c := s[i]
		ls[c]++
		rs[c]--
		if rs[c] == 0 {
			delete(rs, c)
		}
	}

	return cnt
}

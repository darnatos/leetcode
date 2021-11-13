package solution

func FindNumOfValidWords(words []string, puzzles []string) []int {
	m := make(map[int]int, len(words))
	for _, w := range words {
		m[strToMask(w)]++
	}

	res := make([]int, len(puzzles))
	for i, p := range puzzles {
		mask := strToMask(p)

		c, sub, first := 0, mask, 1<<(p[0]-'a')
		for {
			if _, ok := m[sub]; ok && (sub&first) == first {
				c += m[sub]
			}
			if sub == 0 {
				break
			}
			sub = (sub - 1) & mask
		}

		res[i] = c
	}

	return res
}

func strToMask(s string) int {
	mask := 0
	for i := range s {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

package solution

func CharacterReplacement(s string, k int) int {
	counts := [26]int{}

	l, r, c, res := 0, 0, 0, 0
	for r < len(s) {
		counts[s[r]-'A']++
		c = max(c, counts[s[r]-'A'])

		for r-l+1 > c+k {
			counts[s[l]-'A']--
			l++
		}
		res = max(res, r-l+1)
		r++
	}

	return res
}

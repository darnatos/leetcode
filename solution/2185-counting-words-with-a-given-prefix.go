package solution

func prefixCount(words []string, pref string) int {
	res := 0
	for i := range words {
		k := 0
		for j := range words[i] {
			if k < len(pref) && words[i][j] == pref[k] {
				k++
			} else {
				break
			}
			if k == len(pref) {
				res++
				break
			}
		}
	}
	return res
}

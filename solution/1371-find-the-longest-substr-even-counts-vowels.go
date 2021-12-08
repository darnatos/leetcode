package solution

func FindTheLongestSubstring(s string) int {
	var cM = [26]int{1, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0}
	res := -1
	mask := 0
	m := [32]int{}
	for i := range m {
		m[i] = -1
	}
	for i := 0; i < len(s); i++ {
		mask ^= cM[s[i]-'a']
		if mask != 0 && m[mask] == -1 {
			m[mask] = i
		}

		res = max(res, i-m[mask])
	}
	return res
}

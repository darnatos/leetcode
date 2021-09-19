package solution

func CountSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += countPalindromic(s, i, i)
		res += countPalindromic(s, i, i+1)
	}
	return res
}

func countPalindromic(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		count++
		l--
		r++
	}
	return count
}

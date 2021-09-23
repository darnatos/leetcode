package solution

func BreakPalindrome(palindrome string) string {
	p := []rune(palindrome)
	if len(p) <= 1 {
		return ""
	}

	for i := 0; i < len(p)/2; i++ {
		if p[i] != 'a' {
			p[i] = 'a'
			return string(p)
		}
	}

	p[len(p)-1] = 'b'

	return string(p)
}

package solution

func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}

		var sl, sr byte

		if s[l] >= 'a' {
			sl = s[l] - 'a'
		} else {
			sl = s[l] - 'A'
		}

		if s[r] >= 'a' {
			sr = s[r] - 'a'
		} else {
			sr = s[r] - 'A'
		}

		if sl != sr {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

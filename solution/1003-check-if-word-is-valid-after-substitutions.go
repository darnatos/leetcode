package solution

func IsValid(s string) bool {
	if len(s)%3 != 0 {
		return false
	}

	stack := make([]rune, len(s))
	j := -1
	for _, c := range s {
		if c == 'a' || c == 'b' {
			j++
			stack[j] = c
		} else {
			if j == -1 {
				return false
			}
			if stack[j] != 'b' {
				return false
			}
			j--
			if j == -1 {
				return false
			}
			if stack[j] != 'a' {
				return false
			}
			j--
		}
	}
	return j == -1
}

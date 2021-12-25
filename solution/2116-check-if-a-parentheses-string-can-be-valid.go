package solution

func canBeValid(s string, locked string) bool {
	if len(s)&1 == 1 {
		return false
	}

	count := 0
	for i := range s {
		if locked[i] == '1' && s[i] == ')' {
			count--
		} else {
			count++
		}
		if count < 0 {
			return false
		}
	}

	count = 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '1' && s[i] == '(' {
			count--
		} else {
			count++
		}
		if count < 0 {
			return false
		}
	}
	return true
}

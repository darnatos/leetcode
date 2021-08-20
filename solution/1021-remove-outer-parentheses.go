package solution

func RemoveOuterParentheses(s string) string {
	stack := 0
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack += 1
			if stack > 1 {
				res = append(res, s[i])
			}
		} else {
			if stack > 1 {
				res = append(res, s[i])
			}
			stack -= 1
		}
	}
	return string(res)
}

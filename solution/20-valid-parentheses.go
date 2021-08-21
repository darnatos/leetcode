package solution

func IsValidParentheses(s string) bool {
	n := len(s)
	j := 0
	stack := make([]byte, 0, n/2)
	for i := 0; i < n; i++ {
		switch v := s[i]; v {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			j++
			if j > n/2 {
				return false
			}
			stack = append(stack, v)
		case ')':
			if j == 0 || stack[j-1] != '(' {
				return false
			}
			j--
			stack = stack[:j]
		case '}':
			if j == 0 || stack[j-1] != '{' {
				return false
			}
			j--
			stack = stack[:j]
		case ']':
			if j == 0 || stack[j-1] != '[' {
				return false
			}
			j--
			stack = stack[:j]
		default:
			return false

		}
	}
	return j == 0
}

package solution

func LongestValidParentheses(s string) int {
	stack := make([]int, 0, len(s)/2)
	m, left := 0, -1

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		if len(stack) == 0 {
			left = i
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			m = max(m, i-left)
			continue
		}

		m = max(m, i-stack[len(stack)-1])
	}

	return m
}

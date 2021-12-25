package solution

func Calculate2(s string) int {
	stack := make([]int, 0)
	sign := byte('+')
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			cur = 10*cur + int(s[i]-'0')
		}
		if !(s[i] >= '0' && s[i] <= '9') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, cur)
			case '-':
				stack = append(stack, -cur)
			case '*':
				stack[len(stack)-1] *= cur
			case '/':
				stack[len(stack)-1] /= cur
			}
			cur = 0
			sign = s[i]
		}
	}

	sum := 0
	for i := range stack {
		sum += stack[i]
	}
	return sum
}

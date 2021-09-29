package solution

func Calculate(s string) int {
	stack := make([]int, 0, 4)
	sum := 0
	cur := 0
	sign := 1

	for i := range s {
		switch s[i] {
		case ' ':
			break
		case '+':
			sum += sign * cur
			cur = 0
			sign = 1
		case '-':
			sum += sign * cur
			cur = 0
			sign = -1
		case '(':
			stack = append(stack, sum)
			stack = append(stack, sign)
			sum = 0
			sign = 1
		case ')':
			sum += sign * cur
			cur = 0
			sum *= stack[len(stack)-1]
			sum += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		default:
			if isDigit(s[i]) {
				cur *= 10
				cur += int(s[i] - '0')
			}
		}
	}
	if cur != 0 {
		sum += sign * cur
	}
	return sum
}

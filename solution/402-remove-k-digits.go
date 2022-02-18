package solution

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num)-k)
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

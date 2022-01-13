package solution

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	n := len(s)
	n_letters := 0
	for i := range s {
		if letter == s[i] {
			n_letters++
		}
	}
	stack := make([]byte, 0, k)
	for i := range s {
		for len(stack) != 0 && stack[len(stack)-1] > s[i] && n-i+len(stack) > k && (stack[len(stack)-1] != letter || n_letters > repetition) {
			if stack[len(stack)-1] == letter {
				repetition++
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			if s[i] == letter {
				stack = append(stack, s[i])
				repetition--
			} else if k-len(stack) > repetition {
				stack = append(stack, s[i])
			}
		}
		if letter == s[i] {
			n_letters--
		}
	}
	return string(stack)
}

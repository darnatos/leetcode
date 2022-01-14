package solution

import "strings"

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	for i := range s {
		count[s[i]-'a']++
	}
	visited := make([]bool, 26)
	stack := make([]byte, 0, 26)
	for i := range s {
		ch := s[i] - 'a'
		count[ch]--
		if visited[ch] {
			continue
		}
		for len(stack) > 0 && ch < stack[len(stack)-1] && count[stack[len(stack)-1]] > 0 {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		visited[ch] = true
		stack = append(stack, ch)
	}
	var res strings.Builder
	res.Grow(len(stack))
	for i := range stack {
		res.WriteByte(stack[i] + 'a')
	}
	return res.String()
}

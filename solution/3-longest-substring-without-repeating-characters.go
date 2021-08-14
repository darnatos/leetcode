package solution

import "strings"

func LengthOfLongestSubstring(s string) int {
	res := 0
	x := 0

	for i := 0; i < len(s); i++ {
		j := strings.LastIndexByte(s[x:i], s[i])

		if j > -1 {
			x = x + j + 1
		}
		if res < i+1-x {
			res = i + 1 - x
		}
	}
	return res
}

package solution

import (
	"strings"
)

func DecodeString(s string) string {
	res, _ := helper(s, 0)
	return res
}

func helper(s string, i int) (string, int) {
	var sb strings.Builder

	for i < len(s) && s[i] != ']' {
		if !isDigit(s[i]) {
			sb.WriteByte(s[i])
			i++
		} else {
			var n int
			for isDigit(s[i]) {
				n = 10*n + int(s[i]-'0')
				i++
			}
			i++
			var t string
			t, i = helper(s, i)
			i++

			for ; n > 0; n-- {
				sb.WriteString(t)
			}
		}
	}

	return sb.String(), i
}

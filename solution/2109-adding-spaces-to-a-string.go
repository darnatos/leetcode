package solution

import "strings"

func addSpaces(s string, spaces []int) string {
	j := 0
	var res strings.Builder
	res.Grow(len(s) + len(spaces))
	for i := 0; i < len(s); i++ {
		if j < len(spaces) && i >= spaces[j] {
			res.WriteByte(' ')
			j++
		}
		res.WriteByte(s[i])
	}
	return res.String()
}

package solution

import "strconv"

func ThousandSeparator(n int) string {
	s := strconv.Itoa(n)
	for i := len(s)-3; i>0;i-=3{
		s = s[0:i] + "." + s[i:]
	}
	return s
}

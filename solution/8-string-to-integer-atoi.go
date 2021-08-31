package solution

import "math"

func MyAtoi(s string) int {
	i, sign, res := 0, 1, 0

	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else {
			break
		}
	}
	if len(s) <= i {
		return res
	}
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		} else {
			if math.MaxInt32/10 < res || math.MaxInt32/10 == res && s[i]-'0' > 7 {
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			} else {
				res = res*10 + int(s[i]) - '0'
			}
		}
	}
	return sign * res
}

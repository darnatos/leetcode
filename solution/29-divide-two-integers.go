package solution

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}

	if dividend < divisor {
		return 0
	}

	tmp, q := 0, 0
	for i := 31; i >= 0; i-- {
		if tmp+(divisor<<i) <= dividend {
			tmp += divisor << i
			q |= 1 << i
		}
	}

	return sign * q
}

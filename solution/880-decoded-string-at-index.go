package solution

func DecodeAtIndex(s string, k int) string {
	i := 0
	var n int64 = 0
	for ; n < int64(k); i++ {
		if isDigit(s[i]) {
			n *= int64(s[i] - '0')
		} else {
			n++
		}
	}
	for i--; i > 0; i-- {
		if isDigit(s[i]) {
			n /= int64(s[i] - '0')
			k %= int(n)
		} else {
			if int64(k)%n == 0 {
				break
			}
			n--
		}
	}
	return s[i : i+1]
}

package solution

func ShortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	j := -len(s)
	for i := range s {
		if s[i] == c {
			j = i
		}
		res[i] = i - j
	}
	for i := j - 1; i >= 0; i-- {
		if s[i] == c {
			j = i
		}
		if res[i] > j-i {
			res[i] = j - i
		}
	}

	return res
}

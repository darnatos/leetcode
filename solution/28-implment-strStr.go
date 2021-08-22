package solution

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	n := len(needle)
	for i := 0; i < len(haystack)-n+1; i++ {
		if needle == haystack[i:i+n] {
			return i
		}
	}

	return -1
}

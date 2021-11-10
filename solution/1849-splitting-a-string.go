package solution

func SplitString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if ssDfs(s, toVal(s[:i+1]), i+1) {
			return true
		}
	}
	return false
}

func ssDfs(s string, prev, i int) bool {
	if i == len(s) {
		return true
	}
	for j := i + 1; j <= len(s); j++ {
		cur := toVal(s[i:j])
		if cur == prev-1 && ssDfs(s, cur, j) {
			return true
		}
	}
	return false
}

func toVal(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if res == 0 && s[i] == '0' {
			continue
		}
		res *= 10
		res += int(s[i] - '0')
	}
	return res
}

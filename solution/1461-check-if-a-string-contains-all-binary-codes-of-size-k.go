package solution

func hasAllCodes(s string, k int) bool {
	m := make(map[int]struct{})
	mask := 0
	cur := 0
	for i := 0; i < len(s) && i < k; i++ {
		mask = mask<<1 | 1
		cur = cur<<1 + int(s[i]) - '0'
	}
	m[cur] = struct{}{}
	for i := k; i < len(s); i++ {
		cur = mask & (cur<<1 + int(s[i]) - '0')
		m[cur] = struct{}{}
	}
	return len(m) == 1<<k
}

package solution

func GenerateParenthesis(n int) []string {
	return helper(n)
}

func helper(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	s := helper(n-1)

	m := make(map[string]bool, len(s))
	for p := range s {
		for i := range s[p] {
			tmp := s[p][0:i]+"()"+s[p][i:]
			m[tmp] = true
		}
	}

	u := make([]string, 0,len(m))
	for k := range m {
		u = append(u, k)
	}

	return u
}

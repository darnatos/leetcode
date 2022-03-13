package solution

func cellsInRange(s string) []string {
	res := make([]string, 0)
	for j := s[0]; j <= s[3]; j++ {
		for i := s[1]; i <= s[4]; i++ {
			res = append(res, string(j)+string(i))
		}
	}
	return res
}

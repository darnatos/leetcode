package solution

func kSimilarity(s1 string, s2 string) int {
	//bfs
	n := len(s1)
	vis := make(map[string]struct{})
	q := make([]string, 0)
	q = append(q, s1)
	vis[s1] = struct{}{}
	step := 0
	for len(q) > 0 {
		end := len(q)
		for i := 0; i < end; i++ {
			cur := q[i]
			if cur == s2 {
				return step
			}
			s := []byte(cur)
			j0 := 0
			for s[j0] == s2[j0] {
				j0++
			}
			for j := j0 + 1; j < n; j++ {
				if s[j] != s2[j0] || s[j] == s2[j] {
					continue
				}
				s[j0], s[j] = s[j], s[j0]
				ss := string(s)
				if _, ok := vis[ss]; !ok {
					vis[ss] = struct{}{}
					q = append(q, ss)
				}
				s[j0], s[j] = s[j], s[j0]
			}
		}
		q = q[end:]
		step++
	}
	return -1
}

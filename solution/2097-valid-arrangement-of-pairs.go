package solution

func ValidArrangement(pairs [][]int) [][]int {
	s := make(map[int][]int, len(pairs))
	cnt := make(map[int]int, len(pairs))
	for _, p := range pairs {
		s[p[0]] = append(s[p[0]], p[1])
		cnt[p[0]]++
		cnt[p[1]]--
	}
	cur := pairs[0][0]
	for i := range cnt {
		if cnt[i] == 1 {
			cur = i
			break
		}
	}

	route := make([]int, 0, len(pairs)+1)

	vaDfs(s, cur, &route)

	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	res := make([][]int, len(pairs))
	for i := 0; i < len(res); i++ {
		res[i] = []int{route[i], route[i+1]}
	}
	return res
}

func vaDfs(s map[int][]int, cur int, res *[]int) {
	for len(s[cur]) != 0 {
		next := s[cur][0]
		s[cur] = s[cur][1:]
		vaDfs(s, next, res)
	}
	*res = append(*res, cur)
}

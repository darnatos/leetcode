package solution

func isTransformable(s string, t string) bool {
	idx := [10][]int{}
	pos := [10]int{}
	var p *[]int
	for i := range s {
		p = &idx[s[i]-'0']
		*p = append(*p, i)
	}
	for j := range t {
		d := int(t[j] - '0')

		if pos[d] >= len(idx[d]) {
			return false
		}
		for i := 0; i < d; i++ {
			if pos[i] < len(idx[i]) && idx[i][pos[i]] < idx[d][pos[d]] {
				return false
			}
		}

		pos[d]++
	}
	return true
}

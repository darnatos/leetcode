package solution

func countPoints(rings string) int {
	m := make([]int, 20)
	for i := 0; i < len(rings); i += 2 {
		c := 0
		switch rings[i] {
		case 'R':
			c = 1
		case 'G':
			c = 2
		case 'B':
			c = 4
		default:
		}
		m[int(rings[i+1]-'0')] |= c
	}
	res := 0
	for i := range m {
		if m[i] == 7 {
			res++
		}
	}
	return res
}

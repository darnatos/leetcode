package solution

func numberOfWays(corridor string) int {
	pos := make([]int, 0)
	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			pos = append(pos, i)
		}
	}
	if len(pos) == 0 || len(pos)&1 == 1 {
		return 0
	}
	res := 1
	for i := 2; i < len(pos); i += 2 {
		res = (res * (pos[i] - pos[i-1])) % 1000000007
	}
	return res
}

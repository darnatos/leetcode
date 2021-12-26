package solution

func executeInstructions(n int, startPos []int, s string) []int {
	res := make([]int, len(s))
	dir := map[byte][]int{
		'U': {-1, 0},
		'D': {1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}
	for i := 0; i < len(s); i++ {
		count := 0
		x, y := startPos[0], startPos[1]
		for j := i; j < len(s); j++ {
			x, y = x+dir[s[j]][0], y+dir[s[j]][1]
			if x >= 0 && x < n && y >= 0 && y < n {
				count++
			} else {
				break
			}
		}
		res[i] = count
	}
	return res
}

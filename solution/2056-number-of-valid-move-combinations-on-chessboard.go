package solution

func MountCombinations(pieces []string, pos [][]int) int {
	dir := [8][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	poss := map[string][][2]int{
		"rook":   dir[:4],
		"queen":  dir[:],
		"bishop": dir[4:],
	}
	var b [4][8][8]int

	var helper func([]string, [][]int, int) int
	helper = func(pieces []string, pos [][]int, p int) int {
		if p >= len(pieces) {
			return 1
		}
		i, j, res := pos[p][0]-1, pos[p][1]-1, 0
		for _, d := range poss[pieces[p]] {
			blocked := false
			step := 2
			if res == 0 {
				step = 1
			}
			for !blocked {
				x, y := i+(step-1)*d[0], j+(step-1)*d[1]
				if x < 0 || y < 0 || x > 7 || y > 7 {
					break
				}
				canStop := true
				for pp := 0; pp < p; pp++ {
					canStop = canStop && b[pp][x][y] >= 0 && b[pp][x][y] < step
					blocked = blocked || (b[pp][x][y] < 0 && -b[pp][x][y] <= step) || b[pp][x][y] == step
				}
				if canStop {
					b[p][x][y] = -step
					res += helper(pieces, pos, p+1)
				}
				b[p][x][y] = step
				step++
			}
			for m := 0; m < 8; m++ {
				for n := 0; n < 8; n++ {
					b[p][m][n] = 0
				}
			}
		}
		return res
	}
	return helper(pieces, pos, 0)
}

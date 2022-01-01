package solution

import "math"

func minFallingPathSum(grid [][]int) int {
	fm2, sm2, p2 := 0, 0, -1
	for i := 0; i < len(grid); i++ {
		fm, sm, p := math.MaxInt32, math.MaxInt32, -1
		for j := range grid[0] {
			mn := fm2
			if j == p2 {
				mn = sm2
			}
			if grid[i][j]+mn < fm {
				sm = fm
				fm = mn + grid[i][j]
				p = j
			} else if grid[i][j]+mn < sm {
				sm = grid[i][j] + mn
			}
		}
		fm2, sm2, p2 = fm, sm, p
	}

	return fm2
}

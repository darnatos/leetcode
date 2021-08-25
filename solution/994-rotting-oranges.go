package solution

func OrangesRotting(grid [][]int) int {
	adjVector := [...]int{0, -1, 0, 1, 0}
	q := make([][]int, 0, 8)

	c := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				c++
			}
		}
	}
	if c == 0 {
		return 0
	}
	ans := -1
	for len(q) != 0 {
		s := len(q)
		ans++
		for i := 0; i < s; i++ {
			h := q[0]
			q = q[1:]

			for i := 0; i < 4; i++ {
				x := h[0] + adjVector[i]
				y := h[1] + adjVector[i+1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					c--
					q = append(q, []int{x, y})
				}
			}
		}
	}

	if c == 0 {
		return ans
	}

	return -1
}

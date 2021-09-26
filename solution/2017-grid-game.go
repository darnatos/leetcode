package solution

func GridGame(grid [][]int) int64 {
	var a, b int64

	for i := 1; i < len(grid[0]); i++ {
		a += int64(grid[0][i])
	}
	res := a
	for i := 1; i < len(grid[0]); i++ {
		a -= int64(grid[0][i])
		b += int64(grid[1][i-1])
		c := a
		if a < b {
			c = b
		}
		if res > c {
			res = c
		}
	}
	return res
}

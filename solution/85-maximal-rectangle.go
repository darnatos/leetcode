package solution

func MaximalRectangle(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	h := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				h[j] = 0
				continue
			}
			h[j]++
			for cur, pre := j-1, h[j]; cur >= 0; cur-- {
				if h[cur] == 0 {
					break
				}
				pre = min(pre, h[cur])
				res = max(res, pre*(j-cur+1))
			}
			res = max(res, h[j])
		}
	}
	return res
}

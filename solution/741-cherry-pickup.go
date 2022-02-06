package solution

func cherryPickup1(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -2500
			}
		}
	}
	dp[1][1][1] = grid[0][0]
	for x1 := 1; x1 <= n; x1++ {
		for y1 := 1; y1 <= n; y1++ {
			for x2 := 1; x2 <= n; x2++ {
				y2 := x1 + y1 - x2

				if dp[x1][y1][x2] > 0 || y2 < 1 || y2 > n || grid[x1-1][y1-1] == -1 || grid[x2-1][y2-1] == -1 {
					continue
				}
				cur := max(max(dp[x1-1][y1][x2], dp[x1-1][y1][x2-1]), max(dp[x1][y1-1][x2], dp[x1][y1-1][x2-1]))
				if cur < 0 {
					continue
				}
				dp[x1][y1][x2] = cur + grid[x1-1][y1-1]
				if x1 != x2 {
					dp[x1][y1][x2] += grid[x2-1][y2-1]
				}
			}
		}
	}
	//        fmt.Println(dp)

	return max(dp[n][n][n], 0)
}

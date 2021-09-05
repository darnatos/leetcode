package solution

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp, dp0 := make([]int, n+1), make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = math.MaxInt32
		dp0[i] = math.MaxInt32
	}

	dp0[n] = 0
	dp[n-1] = 0
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			dp0[j] = max(min(dp[j], dp0[j+1])-dungeon[i][j],0)
		}
		dp, dp0 = dp0, dp
		dp0[n] = math.MaxInt32
	}
	return 1+dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

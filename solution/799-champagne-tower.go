package solution

func champagneTower(poured int, query_row int, query_glass int) float64 {
	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]float64, query_row+1)
	var rest float64
	dp[0] = []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		dp[i] = make([]float64, i+2)
		for j := 0; j < i; j++ {
			rest = max(dp[i-1][j]-1.0, 0)
			dp[i][j] += rest / 2.0
			dp[i][j+1] += rest / 2.0
		}
	}
	return -max(-dp[query_row][query_glass], -1.0)
}

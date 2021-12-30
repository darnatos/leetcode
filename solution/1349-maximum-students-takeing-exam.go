package solution

import "math/bits"

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	validRows := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seats[i][j] == '.' {
				validRows[i] = validRows[i] | (1 << j)
			}
		}
	}
	stateSize := 1 << n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, stateSize)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < stateSize; j++ {
			if j&validRows[i] == j && j&(j>>1) == 0 {
				if i == 0 {
					dp[i][j] = bits.OnesCount(uint(j))
				} else {
					for k := 0; k < stateSize; k++ {
						if k>>1&j == 0 && j>>1&k == 0 && dp[i-1][k] != -1 {
							dp[i][j] = max(dp[i][j], dp[i-1][k]+bits.OnesCount(uint(j)))
						}
					}
				}
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

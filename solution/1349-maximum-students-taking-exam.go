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

func maxStudentsHungarian(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	dir := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 1}, {0, 1}, {1, 1}}

	matching := make([][][]int, m)
	for i := range matching {
		matching[i] = make([][]int, n)
	}

	var dfs func(node []int, visited [][]bool) bool
	dfs = func(node []int, visited [][]bool) bool {
		i, j := node[0], node[1]
		for _, d := range dir {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] && seats[ni][nj] == '.' {
				visited[ni][nj] = true
				if len(matching[ni][nj]) == 0 || dfs(matching[ni][nj], visited) {
					matching[ni][nj] = []int{i, j}
					return true
				}
			}
		}
		return false
	}

	var Hungarian func() int
	Hungarian = func() int {
		res := 0
		for j := 0; j < n; j += 2 {
			for i := 0; i < m; i++ {
				if seats[i][j] == '.' {
					visited := make([][]bool, m)
					for k := range visited {
						visited[k] = make([]bool, n)
					}
					visited[i][j] = true
					if dfs([]int{i, j}, visited) {
						res++
					}
				}
			}
		}
		return res
	}

	res := Hungarian()

	count := 0
	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == '.' {
				count++
			}
		}
	}
	return count - res
}

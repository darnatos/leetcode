package solution

import "fmt"

func possibleToStamp(grid [][]int, h int, w int) bool {
	m, n := len(grid), len(grid[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i := range grid {
		for j := range grid[0] {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + (1 - grid[i][j])
		}
	}

	stamp := make([][]int, m)
	for i := range stamp {
		stamp[i] = make([]int, n)
	}

	for i := h - 1; i < m; i++ {
		for j := w - 1; j < n; j++ {
			fmt.Println(i, j, sum2d(pre, i, j, i-h+1, j-w+1))
			stamp[i][j] = sum2d(pre, i, j, i-h+1, j-w+1) / w / h
		}
	}

	pre2 := make([][]int, m+1)
	for i := range pre {
		pre2[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pre2[i+1][j+1] = pre2[i+1][j] + pre2[i][j+1] - pre2[i][j] + stamp[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && sum2d(pre2, min(m-1, i+h-1), min(n-1, j+w-1), i, j) == 0 {
				return false
			}
		}
	}
	return true
}
func sum2d(pre [][]int, r1, c1, r2, c2 int) int {
	return pre[r1+1][c1+1] - pre[r1+1][c2] - pre[r2][c1+1] + pre[r2][c2]
}

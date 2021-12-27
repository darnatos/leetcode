package solution

import "math"

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, sessionTime+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(cur, mask int) int
	helper = func(cur, mask int) int {
		if cur > sessionTime {
			return math.MaxInt32
		}
		if mask == 1<<n-1 {
			return 1
		}
		if dp[mask][cur] != -1 {
			return dp[mask][cur]
		}

		res := math.MaxInt32
		for i := range tasks {
			if 1<<i&mask == 0 {
				included := helper(cur+tasks[i], mask|1<<i)
				next := 1 + helper(tasks[i], mask|1<<i)
				res = min(res, min(included, next))
			}
		}
		dp[mask][cur] = res
		return res
	}

	return helper(0, 0)
}

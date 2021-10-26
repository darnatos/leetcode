package solution

import "sort"

func MaxEnvelopes(env [][]int) int {
	sort.Slice(env, func(i, j int) bool {
		if env[i][0] == env[j][0] {
			return env[i][1] > env[j][1]
		}
		return env[i][0] < env[j][0]
	})

	dp := make([]int, len(env))
	res := 0
	for _, e := range env {
		l, r := 0, res
		for l < r {
			m := l + (r-l)/2
			if dp[m] < e[1] {
				l = m + 1
			} else {
				r = m
			}
		}
		dp[l] = e[1]
		if l == res {
			res++
		}
	}

	return res
}

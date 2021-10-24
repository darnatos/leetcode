package solution

import "leetcode/util"

func MinimumTime(n int, relations [][]int, time []int) int {
	adj := make([][]int, n+1)
	for _, r := range relations {
		adj[r[1]] = append(adj[r[1]], r[0])
	}
	cache := make([]int, n+1)
	var dfs func(int)
	dfs = func(x int) {
		if cache[x] > 0 {
			return
		}
		cur := time[x-1]
		tmp := 0
		for _, i := range adj[x] {
			dfs(i)
			if tmp < cache[i] {
				tmp = cache[i]
			}
		}
		cur += tmp
		cache[x] = cur
	}
	res := 0
	for i := 1; i <= n; i++ {
		dfs(i)
		res = util.Max(res, cache[i])
	}

	return res
}

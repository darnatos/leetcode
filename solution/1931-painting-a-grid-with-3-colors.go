package solution

func ColorTheGrid(m int, n int) int {
	dp := make([][]int, n)
	M := 1 << (2 * m)
	tmp := make([]int, n*M)
	for i := range dp {
		dp[i] = tmp[i*M : (i+1)*M]
	}
	var dfs func(m, n, i, j, cur, prev int) int
	dfs = func(m, n, i, j, cur, prev int) int {
		if i == m {
			return dfs(m, n, 0, j+1, 0, cur)
		}
		if j == n {
			return 1
		}
		if i == 0 && dp[j][prev] > 0 {
			return dp[j][prev]
		}
		res, up, left := 0, 0, (prev>>(i*2))&3
		if i != 0 {
			up = (cur >> ((i - 1) * 2)) & 3
		}
		for k := 1; k <= 3; k++ {
			if k != left && k != up {
				res = (res + dfs(m, n, i+1, j, cur+(k<<(i*2)), prev)) % 1000000007
			}
		}
		if i == 0 {
			dp[j][prev] = res
		}
		return res
	}

	return dfs(m, n, 0, 0, 0, 0)
}

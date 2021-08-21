package solution

const mod = 1e9 + 7

func NumRollsToTargetRecur(d int, f int, target int) int {
	memo := make([][]int, d)
	for i := 0; i < d; i++ {
		memo[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memo[i][j] = -1
		}
	}

	return recurNumRollsToTarget(d, f, target, memo)
}

func recurNumRollsToTarget(d, f, target int, memo [][]int) int {
	if d == 0 && target == 0 {
		return 1
	}
	if d == 0 || target < 0 {
		return 0
	}
	if memo[d-1][target] != -1 {
		return memo[d-1][target]
	}
	count := 0
	for face := 1; face <= f; face++ {
		count = (count + recurNumRollsToTarget(d-1, f, target-face, memo)) % mod
	}
	memo[d-1][target] = count
	return count
}

func NumRollsToTargetDP(d int, f int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	tmp := make([]int, target+1)
	for i := 1; i <= d; i++ {
		copy(tmp, dp)
		if i > 1 {
			tmp[0] = 0
		}
		for j := 1; j <= target; j++ {
			count := 0
			for face := 1; face <= f; face++ {
				t := j - face
				if t < 0 {
					break
				}
				count = (count + tmp[t]) % mod
			}
			dp[j] = count
		}
	}

	return dp[target]
}

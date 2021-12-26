package solution

func numberOfArrays(s string, k int) int {
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = -1
	}

	return noaDfs(s, k, 0, dp)
}

func noaDfs(s string, k, from int, dp []int) int {
	if from >= len(s) {
		return 1
	}
	if s[from] == '0' {
		return 0
	}
	if dp[from] >= 0 {
		return dp[from]
	}

	res := 0
	num := 0

	for i := from; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
		if num > k {
			break
		}
		res += noaDfs(s, k, i+1, dp)
		res %= 1000000007
	}

	dp[from] = res
	return res
}

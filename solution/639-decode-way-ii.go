package solution

func NnumDecodings2(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1

	if s[0] == '*' {
		dp[1] = 9
	} else if s[0] >= '1' {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		c1, c2 := s[i-2], s[i-1]

		if c2 == '*' {
			dp[i] += 9 * dp[i-1]
		} else if c2 >= '1' {
			dp[i] += dp[i-1]
		}

		if c1 == '*' {
			if c2 == '*' {
				dp[i] += 15 * dp[i-2]
			} else if c2 <= '6' {
				dp[i] += 2 * dp[i-2]
			} else {
				dp[i] += dp[i-2]
			}
		} else if c1 == '1' {
			if c2 == '*' {
				dp[i] += 9 * dp[i-2]
			} else if c2 <= '9' {
				dp[i] += dp[i-2]
			}
		} else if c1 == '2' {
			if c2 == '*' {
				dp[i] += 6 * dp[i-2]
			} else if c2 <= '6' {
				dp[i] += dp[i-2]
			}
		}
		dp[i] %= mod
	}
	return dp[n]
}

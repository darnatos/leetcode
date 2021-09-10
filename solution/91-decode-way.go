package solution

import "strconv"

func NumDecodings(s string) int {
	dp := make(map[string]int)

	return numDecodings(s, dp)
}

func numDecodings(s string, dp map[string]int) int {
	if v, ok := dp[s]; ok {
		return v
	}
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}

	res := 0
	if len(s) >= 1 {
		res += numDecodings(s[1:], dp)
	}
	if len(s) >= 2 {
		i, err := strconv.Atoi(s[:2])
		if err == nil && i >= 1 && i <= 26 {
			res += numDecodings(s[2:], dp)
		}
	}

	dp[s] = res
	return res
}

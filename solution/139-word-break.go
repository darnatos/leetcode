package solution

func WordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		m[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			if !dp[i] {
				continue
			}
			if !m[s[i:j]] {
				continue
			}
			dp[j] = true
		}
	}

	return dp[len(s)]
}

package solution

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for j := 1; j <= len(s2); j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len(s1); i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= len(s2); j++ {
			dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1] || dp[j] && s1[i-1] == s3[i+j-1]
		}
	}
	return dp[len(s2)]
}

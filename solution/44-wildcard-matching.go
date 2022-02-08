package solution

func isMatch2(s string, p string) bool {
	//dp[i+1][j+1] denotes the result of isMatch(s[:i+1], p[:j+1])
	//if s[i]==p[j] || '?'==p[j], dp[i+1][j+1]=dp[i][j]
	//if p[j]=='*', dp[i+1][j+1]=dp[i][j+1]||dp[i+1][j]
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	dp0 := make([]bool, n+1)

	// s=="" and p=="" match
	dp[0] = true
	for j := 0; j < n; j++ {
		if p[j] == '*' {
			dp[j+1] = dp[j]
		} else {
			break
		}
	}
	for i := 0; i < m; i++ {
		dp, dp0 = dp0, dp
		// s!="" and p=="" must not match
		dp[0] = false
		for j := range p {
			// if s[i]==p[j] || '?'==p[j]{
			//     dp[j+1]=dp0[j]
			// } else if '*'==p[j]{
			//     dp[j+1]=dp0[j+1]||dp[j]
			// } else {
			//     // reset dp[j+1] to false to be reused as dp0 in the next itr
			//     dp[j+1]=false
			// }
			dp[j+1] = (s[i] == p[j] || '?' == p[j]) && dp0[j] || '*' == p[j] && (dp0[j+1] || dp[j])
		}
	}

	return dp[n]
}

package solution

import "fmt"

func NumPermsDISequence(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp2 := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		cur := 0
		if s[i] == 'I' {
			for j := 0; j < n-i; j++ {
				cur = (cur + dp[j]) % mod
				dp2[j] = cur
			}
		} else {
			cur := 0
			for j := n - 1 - i; j >= 0; j-- {
				cur = (cur + dp[j+1]) % mod
				dp2[j] = cur
			}
		}
		fmt.Println(dp, dp2)
		copy(dp, dp2)
	}
	return dp[0]
}

package solution

import (
	"fmt"
	"math"
)

func NumSquares(n int) int {
	all := make([]int, 0, 101)
	j := 1
	for i := 1; j <= n; i++ {
		j = i * i
		all = append(all, j)
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for _, j := range all {
			if j > i {
				break
			}
			dp[i] = min(dp[i], dp[i-j]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

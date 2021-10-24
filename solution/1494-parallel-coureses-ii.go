package solution

import (
	"leetcode/util"
	"math/bits"
)

func MinNumberOfSemesters(n int, dependencies [][]int, k int) int {
	dp := make([]int, 1<<n)
	preReq := make([]int, 1<<n)
	preCourse := make([]int, n)
	for _, dep := range dependencies {
		dep[0]--
		dep[1]--
		preCourse[dep[1]] += 1 << (dep[0])
	}

	for state := 1; state < 1<<n; state++ {
		dp[state] = n
		for i := 0; i < n; i++ {
			if state>>i&1 == 1 {
				preReq[state] |= preCourse[i]
			}
		}
	}

	for state := 1; state < 1<<n; state++ {
		preState := state
		for preState > 0 {
			preState = (preState - 1) & state
			if preReq[state] == preState&preReq[state] && courseDiff(state, preState) <= k {
				dp[state] = util.Min(dp[state], dp[preState]+1)
			}
		}
	}

	return dp[1<<n-1]
}

func courseDiff(a, b int) int {
	return bits.OnesCount(uint(a)) - bits.OnesCount(uint(b))
}

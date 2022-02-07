package solution

func superEggDrop(k int, n int) int {
	dp := make([]int, k+1)

	// dp[k] means if we have k eggs, after m moves,
	// max floors can be covered.
	// if dp[k] >= k, then m moves is the minimum
	// number of moves since we increase m by 1 from 0
	// TC: O(k*log(n)), SC: O(k)
	m := 0
	for dp[k] < n {
		m++
		for i := k; i >= 1; i-- {
			dp[i] = dp[i-1] + dp[i] + 1
		}
	}

	return m
}

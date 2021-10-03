package solution

func BBMaxCoins(nums []int) int {
	n := len(nums)

	nums = append(nums, []int{1, 1}...)
	copy(nums[1:], nums)
	nums[0] = 1

	tmp := make([]int, (n+2)*(n+1))
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = tmp[i*(n+1) : (i+1)*(n+1)]
		if i > 0 && i < n+1 {
			dp[i][i] = nums[i-1] * nums[i] * nums[i+1]
		}
	}

	for le := 1; le <= n; le++ {
		for i := 1; i <= n+1-le; i++ {
			j := i + le - 1
			for k := i; k <= j; k++ {
				dp[i][j] = max(dp[i][j], nums[i-1]*nums[k]*nums[j+1]+dp[i][k-1]+dp[k+1][j])
			}
		}
	}

	return dp[1][n]
}

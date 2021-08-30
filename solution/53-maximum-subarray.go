package solution

func MaxSubArray(nums []int) int {
	max := nums[0]

	dp := make([]int, len(nums))
	dp[0] = max

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}

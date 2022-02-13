package solution

import "math/bits"

func maximumANDSum(nums []int, numSlots int) int {
	for len(nums) < numSlots*2 {
		nums = append(nums, 0)
	}
	n := len(nums)
	res := 0
	dp := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		slot := (bits.OnesCount(uint(i)) + 1) / 2
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				dp[i] = max(dp[i], dp[i^(1<<j)]+slot&nums[j])
			}
			// fmt.Printf("%06b %v %v\n",i^(1<<j),slot,dp[i])
		}
		res = max(res, dp[i])
	}

	return res
}

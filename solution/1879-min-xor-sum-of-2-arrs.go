package solution

import "math"

func minimumXORSum(nums1 []int, nums2 []int) int {
	dp := make([]int, 1<<len(nums1))
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	return mxsDfs(nums1, nums2, 0, 0, dp)
}

func mxsDfs(nums1, nums2 []int, i, mask int, dp []int) int {
	if i >= len(nums1) {
		return 0
	}
	if dp[mask] == math.MaxInt32 {
		for j := 0; j < len(nums1); j++ {
			if mask&(1<<j) == 0 {
				dp[mask] = min(dp[mask], (nums1[i]^nums2[j])+mxsDfs(nums1, nums2, i+1, mask+(1<<j), dp))
			}
		}
	}
	return dp[mask]
}

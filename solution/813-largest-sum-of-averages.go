package solution

import "math"

func LargestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	dp := make([]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = float64(sum[n]-sum[i]) / float64(n-i)
	}
	for kk := 1; kk < k; kk++ {
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				dp[i] = math.Max(dp[i], float64(sum[j]-sum[i])/float64(j-i)+dp[j])
			}
		}
	}

	return dp[0]
}

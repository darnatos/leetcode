package solution

import "sort"

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n, maxI := len(nums), 0
	dp, pre := make([]int, n), make([]int, n)
	for i := range dp {
		dp[i] = 1
		pre[i] = -1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				pre[i] = j
			}
		}
		if dp[maxI] < dp[i] {
			maxI = i
		}
	}
	res := make([]int, 0)
	for ; maxI >= 0; maxI = pre[maxI] {
		res = append(res, nums[maxI])
	}
	return res
}

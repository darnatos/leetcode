package solution

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	dp := make([]int, 0, n)
	inc := make([]int, n)
	for i := 0; i < n; i++ {
		l, r := 0, len(dp)
		for l < r {
			m := l + (r-l)/2
			if dp[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[l] = nums[i]
		}
		inc[i] = len(dp)
	}

	dp = dp[:0]
	res := 0
	for i := n - 1; i >= 0; i-- {
		l, r := 0, len(dp)
		for l < r {
			m := l + (r-l)/2
			if dp[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[l] = nums[i]
		}
		if len(dp) > 1 && inc[i] > 1 {
			res = max(res, len(dp)+inc[i])
		}
	}

	return len(nums) - res + 1
}

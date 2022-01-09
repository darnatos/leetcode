package solution

func minSwaps(nums []int) int {
	n := len(nums)
	ones := 0
	for i := range nums {
		ones += nums[i]
	}
	res := ones
	onesInWindow := 0
	for i := n - ones; i < n; i++ {
		onesInWindow += nums[i]
	}

	for i, j := 0, (n-ones)%n; i < n; i, j = i+1, j+1 {
		onesInWindow += nums[i%n] - nums[j%n]
		res = min(res, ones-onesInWindow)
	}
	return res
}

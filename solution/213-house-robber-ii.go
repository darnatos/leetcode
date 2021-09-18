package solution

func Rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	return max(rob2Helper(nums, 0, len(nums)-2), rob2Helper(nums, 1, len(nums)-1))
}

func rob2Helper(nums []int, s, e int) int {
	pp, p := 0, nums[s]
	for i := s + 1; i <= e; i++ {
		p, pp = max(pp+nums[i], p), p
	}
	return p
}

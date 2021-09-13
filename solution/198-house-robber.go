package solution

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pp, p := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		p, pp = max(pp + nums[i], p), p
	}

	return p
}

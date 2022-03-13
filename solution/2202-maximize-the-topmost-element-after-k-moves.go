package solution

func maximumTop(nums []int, k int) int {
	if len(nums) == 1 {
		if k&1 == 1 {
			return -1
		} else {
			return nums[0]
		}
	}
	if k > len(nums) {
		res := -1
		for _, n := range nums {
			if res < n {
				res = n
			}
		}
		return res
	}

	mx := -1
	for i := 0; i < k-1; i++ {
		if mx < nums[i] {
			mx = nums[i]
		}
	}
	if k < len(nums) && nums[k] > mx {
		return nums[k]
	}
	return mx
}

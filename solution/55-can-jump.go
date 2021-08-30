package solution

import "leetcode/util"

func CanJump(nums []int) bool {
	n := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if i > n {
			return false
		}
		n = util.Max(n, i+nums[i])
	}

	return n >= len(nums)-1
}

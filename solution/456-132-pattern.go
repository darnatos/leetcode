package solution

import "math"

func find132pattern(nums []int) bool {
	s3 := math.MinInt32
	// store possible s2
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		// nums[i] is s1
		if nums[i] < s3 {
			return true
		}
		// after the loop, stack is a ascending array
		// and all elements are less than s3
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			s3 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

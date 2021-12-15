package solution

import "math"

func maximumProduct(nums []int) int {
	mx1, mx2, mx3 := math.MinInt32, math.MinInt32, math.MinInt32
	mi1, mi2 := math.MaxInt32, math.MaxInt32
	for i := range nums {
		if nums[i] > mx1 {
			mx1, mx2, mx3 = nums[i], mx1, mx2
		} else if nums[i] > mx2 {
			mx2, mx3 = nums[i], mx2
		} else if nums[i] > mx3 {
			mx3 = nums[i]
		}
		if mi1 > nums[i] {
			mi1, mi2 = nums[i], mi1
		} else if mi2 > nums[i] {
			mi2 = nums[i]
		}
	}
	return max(mx1*mx2*mx3, mi1*mi2*mx1)
}

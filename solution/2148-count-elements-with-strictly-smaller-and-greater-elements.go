package solution

import "sort"

func countElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	l, r := 1, len(nums)-2

	for l <= r && nums[l] == nums[l-1] {
		l++
	}
	for r >= l && nums[r] == nums[r+1] {
		r--
	}
	return r - l + 1
}

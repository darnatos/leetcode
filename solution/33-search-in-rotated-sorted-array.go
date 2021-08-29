package solution

import "fmt"

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		fmt.Println(l, m, r)
		if nums[m] == target {
			return m
		} else if nums[m] < nums[l] && (target < nums[m] || target > nums[l]) {
			r = m - 1
		} else if nums[m] >= nums[l] && (target > nums[m] || target < nums[l]) {
			l = m + 1
		} else if nums[m] < nums[l] {
			l = m + 1
		} else if nums[m] >= nums[l] {
			r = m - 1
		}
	}
	return -1
}

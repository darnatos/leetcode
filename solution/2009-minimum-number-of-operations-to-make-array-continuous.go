package solution

import "sort"

func MinOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	m := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[m] {
			continue
		} else {
			m++
			nums[m] = nums[i]
		}
	}
	nums = nums[0 : m+1]

	res := n
	for i, j := 0, 0; i < len(nums); i++ {
		for ; j < len(nums) && nums[j] < nums[i]+n; j++ {
		}
		res = min(res, n-j+i)
	}
	return res
}

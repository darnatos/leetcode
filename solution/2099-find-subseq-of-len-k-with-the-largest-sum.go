package solution

import "sort"

func maxSubsequence(nums []int, k int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	m := make(map[int]int, len(nums))
	cnt := 0
	for i := len(sorted) - 1; i >= 0; i-- {
		cnt++
		m[sorted[i]]++
		if cnt == k {
			break
		}
	}
	res := make([]int, 0, k)
	for i := range nums {
		if m[nums[i]] > 0 {
			m[nums[i]]--
			res = append(res, nums[i])
		}
	}
	return res
}

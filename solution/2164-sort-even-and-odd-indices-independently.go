package solution

import "sort"

func sortEvenOdd(nums []int) []int {
	e, o := make([]int, 0, len(nums)/2+1), make([]int, 0, len(nums)/2+1)
	for i := range nums {
		if i&1 == 0 {
			e = append(e, nums[i])
		} else {
			o = append(o, nums[i])
		}
	}
	sort.Ints(e)
	sort.Slice(o, func(a, b int) bool { return o[a] > o[b] })

	ee, oo := 0, 0
	for i := range nums {
		if i&1 == 0 {
			nums[i] = e[ee]
			ee++
		} else {
			nums[i] = o[oo]
			oo++
		}
	}
	return nums
}

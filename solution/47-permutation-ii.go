package solution

import "sort"

func PermuteUnique(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Ints(nums)

	permuteBacktrack2(&res, make([]int, 0, len(nums)), make(map[int]bool, len(nums)), nums)
	return res
}

func permuteBacktrack2(res *[][]int, cur []int, cm map[int]bool, nums []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
	} else {
		for i := 0; i < len(nums); i++ {
			if cm[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && !cm[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			cm[i] = true
			permuteBacktrack2(res, cur, cm, nums)
			cur = cur[:len(cur)-1]
			cm[i] = false
		}
	}
}

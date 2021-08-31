package solution

func Permute(nums []int) [][]int {
	res := make([][]int, 0, 4)
	permuteBacktrack(&res, make([]int, 0, len(nums)), make(map[int]bool, len(nums)), nums)
	return res
}

func permuteBacktrack(res *[][]int, cur []int, cm map[int]bool, nums []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
	} else {
		for i := range nums {
			if _, ok := cm[nums[i]]; ok {
				continue
			}
			cur = append(cur, nums[i])
			cm[nums[i]] = true
			permuteBacktrack(res, cur, cm, nums)
			cur = cur[:len(cur)-1]
			delete(cm, nums[i])
		}
	}
}

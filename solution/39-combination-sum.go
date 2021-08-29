package solution

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0, 4)

	combinationSumHelper(candidates, target, []int{}, &res)

	return res
}

func combinationSumHelper(candidates []int, target int, cur []int, res *[][]int) {
	for i, v := range candidates {
		if target <= v {
			if target == v {
				match := make([]int, len(cur)+1)
				copy(match, cur)
				match[len(cur)] = v
				*res = append(*res, match)
			}
			break
		}
		combinationSumHelper(candidates[i:], target-v, append(cur, v), res)
	}
}

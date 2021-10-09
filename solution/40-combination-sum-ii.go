package solution

import (
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	chosen := make([]bool, len(candidates))
	res := make([][]int, 0, 4)

	combinationSum2Helper(candidates, chosen, 0, 0, target, []int{}, &res)

	return res
}

func combinationSum2Helper(candidates []int, chosen []bool, s, sum, target int, cur []int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := s; i < len(candidates); i++ {
		v := candidates[i]
		if i > 0 && candidates[i] == candidates[i-1] && !chosen[i-1] {
			continue
		}
		if target < sum {
			break
		}
		chosen[i] = true
		cur = append(cur, v)
		combinationSum2Helper(candidates, chosen, i+1, sum+v, target, cur, res)
		chosen[i] = false
		cur = cur[:len(cur)-1]
	}
}

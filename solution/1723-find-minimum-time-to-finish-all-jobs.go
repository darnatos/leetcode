package solution

import (
	"math"
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {
	sort.Ints(jobs)
	worker := make([]int, k)
	res := math.MaxInt32
	mtrDfs(jobs, worker, len(jobs)-1, 0, &res)
	return res
}

func mtrDfs(jobs, worker []int, j, cur int, res *int) {
	if cur >= *res {
		return
	}
	if j < 0 {
		*res = cur
		return
	}
	workTime := make(map[int]struct{})
	for i := range worker {
		if _, ok := workTime[worker[i]]; ok {
			continue
		} else {
			workTime[worker[i]] = struct{}{}
		}
		worker[i] += jobs[j]
		mtrDfs(jobs, worker, j-1, max(cur, worker[i]), res)
		worker[i] -= jobs[j]
	}
	return
}

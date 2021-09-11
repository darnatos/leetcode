package solution

import "sort"

type JobPair struct {
	difficulty int
	profit     int
}

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	job := make([]JobPair, len(profit))
	for i := range profit {
		job[i] = JobPair{difficulty[i], profit[i]}
	}
	sort.Slice(job, func(a, b int) bool { return job[a].difficulty < job[b].difficulty })
	sort.Ints(worker)

	res := 0
	j, cur := 0, 0
	for i := range worker {
		for j < len(profit) && job[j].difficulty <= worker[i] {
			if cur < job[j].profit {
				cur = job[j].profit
			}
			j++
		}
		res += cur
	}
	return res
}

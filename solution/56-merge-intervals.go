package solution

import (
	"leetcode/util"
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res := make([][]int, 0, 8)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if len(tmp) > 0 {
			if tmp[1] >= intervals[i][0] {
				tmp[1] = util.Max(tmp[1], intervals[i][1])
				continue
			} else {
				res = append(res, tmp)
			}
		}
		tmp = intervals[i]
	}
	if len(tmp) > 0 {
		res = append(res, tmp)
	}

	return res
}

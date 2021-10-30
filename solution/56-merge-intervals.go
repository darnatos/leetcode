package solution

import (
	"github.com/darnatos/leetcode/util"
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res := make([][]int, 0, 8)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = util.Max(res[len(res)-1][1], intervals[i][1])
		}
	}

	return res
}

package solution

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	i := 0
	res := make([][]int, 0, len(intervals)+1)
	for ;i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ;i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	res = append(res, newInterval)
	res = append(res, intervals[i:]...)
	return res
}

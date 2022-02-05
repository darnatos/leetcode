package solution

import (
	"math"
	"strconv"
)

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	targetTime := make([]string, 0)
	minutes := 0
	for targetSeconds >= 0 {
		secStr := strconv.Itoa(targetSeconds)
		minStr := strconv.Itoa(minutes)
		var sm string
		if minutes == 0 {
			sm = secStr
		} else if len(secStr) == 2 {
			sm = minStr + secStr
		} else {
			sm = minStr + "0" + secStr
		}
		if len(secStr) <= 2 && len(sm) <= 4 {
			targetTime = append(targetTime, sm)
		}
		minutes++
		targetSeconds -= 60
	}

	var dfs func(t string, curAt, i int) int
	dfs = func(t string, curAt, i int) int {
		if len(t) == i {
			return 0
		}
		cur := 0
		if curAt != int(t[i]-'0') {
			cur += moveCost
			curAt = int(t[i] - '0')
		}
		cur += pushCost
		return cur + dfs(t, curAt, i+1)
	}

	res := math.MaxInt32
	for i := range targetTime {
		res = min(res, dfs(targetTime[i], startAt, 0))
	}

	return res
}

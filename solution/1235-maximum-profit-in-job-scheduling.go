package solution

import "sort"

func JobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][]int, len(startTime))
	for i := range jobs {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	dp := make([][]int, 1, len(jobs)+1)
	dp[0] = []int{0, 0}

	for _, job := range jobs {
		cur := dp[floorEntry(dp, job[0])][1] + job[2]
		if cur > dp[len(dp)-1][1] {
			dp = append(dp, []int{job[1], cur})
		}
	}

	return dp[len(dp)-1][1]
}

func floorEntry(dp [][]int, startTime int) int {
	lo, hi := 0, len(dp)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if dp[mid][0] <= startTime {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi - 1
}

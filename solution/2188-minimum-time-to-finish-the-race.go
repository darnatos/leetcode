package solution

import "math"

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	dp := make([]int, numLaps+1)
	for i := 1; i <= numLaps; i++ {
		dp[i] = math.MaxInt32
	}
	for _, t := range tires {
		cur, tot := t[0], t[0]
		for i := 1; i <= numLaps && t[0]+changeTime > cur; i++ {
			dp[i] = min(dp[i], tot)
			cur = cur * t[1]
			tot += cur
		}
	}
	for i := 1; i <= numLaps; i++ {
		for j := 1; j < min(i/2+1, 20); j++ {
			dp[i] = min(dp[i], dp[j]+changeTime+dp[i-j])
		}
	}
	return dp[numLaps]
}

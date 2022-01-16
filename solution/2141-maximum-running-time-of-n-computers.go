package solution

import "sort"

/*
Intuition:

The maximum number of minutes to run n computers simultaneously
must occur when the charge of all batteries, which is `Sum(batteries)`,
could be evenly allocated to all computers and no waste, hence the
upper bound is `Sum(batteries)/n`.


Greedy alg:

When the battery having maximum charge is less than or equals to the current
upper bound, in the other hand, Max(batteries) <= Sum(batteries)/n,
the battery won't get wasted, so do the other smaller.

*/
func maxRunTime(n int, batteries []int) int64 {
	if n > len(batteries) {
		return 0
	}
	sort.Slice(batteries, func(i, j int) bool { return batteries[i] > batteries[j] })

	sum := int64(0)
	for i := range batteries {
		sum += int64(batteries[i])
	}
	for i := range batteries {
		if int64(batteries[i]) > sum/int64(n) {
			sum -= int64(batteries[i])
			n--
		} else {
			return sum / int64(n)
		}
	}
	return 0
}

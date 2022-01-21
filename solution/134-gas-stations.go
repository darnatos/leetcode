package solution

func canCompleteCircuit(gas []int, cost []int) int {
	start, end := len(gas)-1, 0
	sum := gas[start] - cost[start]
	for start > end {
		if sum >= 0 {
			sum += gas[end] - cost[end]
			end++
		} else {
			start--
			sum += gas[start] - cost[start]
		}
	}
	if sum >= 0 {
		return start
	}
	return -1
}

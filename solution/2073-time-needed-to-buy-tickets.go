package solution

func TimeRequiredToBuy(tickets []int, k int) int {
	i := 0
	res := 0
	for tickets[k] > 0 {
		if tickets[i] > 0 {
			tickets[i]--
			res++
		}
		i++
		i %= len(tickets)
	}
	return res
}

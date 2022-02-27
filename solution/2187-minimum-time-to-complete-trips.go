package solution

func minimumTime2(time []int, totalTrips int) int64 {
	var l, r, trips int64 = 0, int64(1e14), 0
	t := int64(totalTrips)
	for l < r {
		m := l + (r-l)/2
		trips = 0
		for i := range time {
			trips += m / int64(time[i])
		}
		if trips >= t {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

package solution

func IsSelfCrossing(d []int) bool {
	for i := 3; i < len(d); i++ {
		if d[i-3] >= d[i-1] && d[i-2] <= d[i] {
			return true
		}
		if i >= 4 {
			if d[i-1] == d[i-3] && d[i-2] <= d[i]+d[i-4] {
				return true
			}
		}
		if i >= 5 {
			if d[i-3]-d[i-5] <= d[i-1] && d[i-1] <= d[i-3] && d[i-2]-d[i-4] <= d[i] && d[i-2]-d[i-4] >= 0 {
				return true
			}
		}
	}

	return false
}

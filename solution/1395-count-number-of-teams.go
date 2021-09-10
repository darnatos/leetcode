package solution

func NumTeams(rating []int) int {
	res := 0
	for i := 1; i < len(rating)-1; i++ {
		ll, lg, rl, rg := 0, 0, 0, 0
		for j := 0; j < len(rating); j++ {
			if i < j {
				if rating[i] < rating[j] {
					lg++
				} else {
					ll++
				}
			} else if i > j {
				if rating[i] < rating[j] {
					rl++
				} else {
					rg++
				}
			}
		}
		res += ll*rl + lg*rg
	}
	return res
}

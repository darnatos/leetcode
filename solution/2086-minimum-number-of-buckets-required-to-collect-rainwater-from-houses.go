package solution

func MinimumBuckets(street string) int {
	if len(street) == 1 {
		if street[0] == 'H' {
			return -1
		} else {
			return 0
		}
	}
	res := 0
	dp := make([]bool, len(street))
	for i := 0; i < len(street); i++ {
		if street[i] == 'H' {
			if i > 0 && dp[i-1] || i+1 < len(street) && dp[i+1] {
				continue
			} else if i+1 < len(street) && street[i+1] == '.' {
				dp[i+1] = true
			} else if i > 0 && street[i-1] == '.' {
				dp[i-1] = true
			} else {
				return -1
			}
		}
	}
	for i := range dp {
		if dp[i] {
			res++
		}
	}
	return res
}

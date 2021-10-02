package solution

func NumOfPairs(nums []string, target string) int {
	res := 0
	m := make(map[string]int, len(nums))
	for i := range nums {
		m[nums[i]]++
	}

	for _, v := range nums {
		if len(v) > len(target) {
			continue
		}
		pref, suff := target[:len(v)], target[len(v):]
		if pref == v {
			res += m[suff]
			if pref == suff {
				res--
			}
		}
	}
	return res
}

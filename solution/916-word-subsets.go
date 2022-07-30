package solution

func wordSubsets(words1 []string, words2 []string) []string {
	cnt := make([]int, 26)
	for _, word := range words2 {
		tmp := make([]int, 26)
		for _, ch := range word {
			tmp[ch-'a']++
		}
		for i := range cnt {
			cnt[i] = max(cnt[i], tmp[i])
		}
	}
	res := make([]string, 0)
	for _, word := range words1 {
		tmp := make([]int, 26)
		for _, ch := range word {
			tmp[ch-'a']++
		}
		flag := true
		for i := range tmp {
			if tmp[i] < cnt[i] {
				flag = false
			}
		}

		if flag {
			res = append(res, word)
		}
	}
	return res
}

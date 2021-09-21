package solution

func WordBreak2(s string, wordDict []string) []string {
	res := make([]string, 0)

	m := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		m[wordDict[i]] = true
	}

	wb2Helper(&res, 0, "", s, m)

	return res
}

func wb2Helper(res *[]string, i int, path, s string, m map[string]bool) {
	if len(s) == i {
		*res = append(*res, path[:len(path)-1])
		return
	}

	for j := i + 1; j <= len(s); j++ {
		if _, ok := m[s[i:j]]; ok {
			wb2Helper(res, j, path+s[i:j]+" ", s, m)
		}
	}
}

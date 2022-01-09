package solution

func wordCount(startWords []string, targetWords []string) int {
	combCount := make(map[int]int, len(targetWords))
	res := make(map[int]int)

	for i := range targetWords {
		mask := 0
		for j := range targetWords[i] {
			mask = mask | (1 << (targetWords[i][j] - 'a'))
		}
		combCount[mask]++
	}

	for i := range startWords {
		mask := 0
		for j := range startWords[i] {
			mask = mask | (1 << (startWords[i][j] - 'a'))
		}
		for j := 0; j < 26; j++ {
			if mask&(1<<j) > 0 {
				continue
			}
			if v, ok := combCount[mask|1<<j]; ok {
				res[mask|1<<j] = v
			}
		}
	}
	sum := 0
	for i := range res {
		sum += res[i]
	}
	return sum
}

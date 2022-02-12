package solution

func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0
	words := make(map[string]bool, len(wordList))
	for i := range wordList {
		words[wordList[i]] = true
	}
	q := make([]string, 0)
	q = append(q, beginWord)
	words[beginWord] = false
	for len(q) > 0 {
		step++
		end := len(q)
		for i := 0; i < end; i++ {
			if q[i] == endWord {
				return step
			}
			cur := []byte(q[i])
			for j := range cur {
				for k := byte('a'); k <= 'z'; k++ {
					cur[j] = k
					tmp := string(cur)
					if words[tmp] {
						words[tmp] = false
						q = append(q, tmp)
					}
				}
				cur[j] = q[i][j]
			}
		}
		q = q[end:]
	}
	return 0
}

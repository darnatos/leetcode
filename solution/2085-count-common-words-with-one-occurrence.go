package solution

func CountWords(words1 []string, words2 []string) int {
	m1, m2 := make(map[string]int, len(words1)), make(map[string]int, len(words2))
	for i := range words1 {
		m1[words1[i]]++
	}
	for i := range words2 {
		m2[words2[i]]++
	}
	res := 0
	for i := range m1 {
		if m1[i] == 1 && m2[i] == 1 {
			res++
		}
	}
	return res
}

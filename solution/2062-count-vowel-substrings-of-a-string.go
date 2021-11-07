package solution

func CountVowelSubstrings(word string) int {
	j, k, res := 0, 0, 0
	count := make(map[byte]int, 5)

	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			count[word[i]]++
			for len(count) == 5 {
				count[word[k]]--
				if count[word[k]] == 0 {
					delete(count, word[k])
				}
				k++
			}
			res += k - j
		} else {
			for j := range count {
				delete(count, j)
			}
			j = i + 1
			k = j
		}
	}

	return res
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

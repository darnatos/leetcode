package solution

func CountVowels(word string) int64 {
	n := len(word)
	var res int64
	for i := 0; i < n; i++ {
		if isVowel(word[i]) {
			res += int64(i+1) * int64(n-i)
		}
	}
	return res
}

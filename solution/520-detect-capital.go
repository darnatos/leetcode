package solution

func detectCapitalUse(word string) bool {
	capCount := 0
	for i := 0; i < len(word); i++ {
		if word[i] <= 'Z' && word[i] >= 'A' {
			capCount++
		}
	}

	return capCount == 0 || capCount == len(word) || word[0] <= 'Z' && word[0] >= 'A' && capCount == 1
}

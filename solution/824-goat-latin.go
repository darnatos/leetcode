package solution

import "strings"

const GoatSuffix = "ma"

func ToGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")

	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	for i := range words {
		if vowels[words[i][0]] {
			words[i] += GoatSuffix + strings.Repeat("a", i+1)
		} else {
			words[i] = words[i][1:] + words[i][:1] + GoatSuffix + strings.Repeat("a", i+1)
		}
	}

	return strings.Join(words, " ")
}

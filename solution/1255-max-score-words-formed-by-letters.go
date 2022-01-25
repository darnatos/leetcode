package solution

func maxScoreWords(words []string, letters []byte, score []int) int {
	cnt := make([]int, 26)
	for i := range letters {
		cnt[letters[i]-'a']++
	}

	return mswBacktrack(words, score, cnt, 0)
}

func mswBacktrack(words []string, score []int, cnt []int, i int) int {
	if i == len(words) {
		return 0
	}

	cur, flag := 0, true
	for _, ch := range words[i] {
		idx := ch - 'a'
		if cnt[idx] == 0 {
			flag = false
		}
		cnt[idx]--
		cur += score[idx]
	}

	if flag {
		cur += mswBacktrack(words, score, cnt, i+1)
	} else {
		cur = 0
	}

	for _, ch := range words[i] {
		cnt[ch-'a']++
	}

	return max(cur, mswBacktrack(words, score, cnt, i+1))
}

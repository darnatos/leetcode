package solution

func LongestSubsequenceRepeatedK(s string, k int) string {
	fre, total := make([]int, 26), 0
	for i := range s {
		fre[s[i]-'a']++
	}
	for i := range fre {
		fre[i] /= k
		total += fre[i]
	}

	candidates := make([]string, 0)
	for l := total; l >= 0; l-- {
		generate(&candidates, fre, 0, l, "")
	}

	for _, cand := range candidates {
		if check(s, cand, k) {
			return cand
		}
	}

	return ""
}

func generate(candidates *[]string, fre []int, cur, l int, seq string) {
	if cur == l {
		*candidates = append(*candidates, seq)
		return
	}

	for i := 25; i >= 0; i-- {
		if fre[i] == 0 {
			continue
		}
		seq += string(rune(i + 'a'))
		fre[i]--
		generate(candidates, fre, cur+1, l, seq)
		seq = seq[:len(seq)-1]
		fre[i]++
	}
}

func check(s, seq string, k int) bool {
	cnt, p := 0, 0
	if len(seq) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == seq[p] {
			p++
		}
		if p == len(seq) {
			cnt++
			p = 0
		}
	}
	return cnt >= k
}

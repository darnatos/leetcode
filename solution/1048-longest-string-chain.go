package solution

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(a int, b int) bool { return len(words[a]) < len(words[b]) })
	cnt := make([]int, len(words))
	res := 0
	for i := range words {
		cnt[i] = lscHelper(words, i, cnt)
		res = max(res, cnt[i])
	}
	return res
}

func lscHelper(words []string, i int, cnt []int) int {
	if cnt[i] > 0 {
		return cnt[i]
	}
	cnt[i] = 1
	for j := 0; j < i; j++ {
		if isPredecessorOf(words[j], words[i]) {
			cnt[i] = max(cnt[i], cnt[j]+1)
		}
	}
	return cnt[i]
}

func isPredecessorOf(a string, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}

package solution

import "math"

func BeautySum(s string) int {
	res := 0
	for l := 3; l <= len(s); l++ {
		freq := make([]int, 26)
		for i := 0; i < l; i++ {
			freq[s[i]-'a']++
		}
		for i := l; i < len(s); i++ {
			res += countDiff(freq)
			freq[s[i-l]-'a']--
			freq[s[i]-'a']++
		}
		res += countDiff(freq)
	}
	return res
}

func countDiff(freq []int) int {
	minCount, maxCount := math.MaxInt32, 0
	for _, c := range freq {
		if c == 0 {
			continue
		}
		if minCount > c {
			minCount = c
		}
		if maxCount < c {
			maxCount = c
		}
	}
	return maxCount - minCount
}

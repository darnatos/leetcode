package solution

import "github.com/darnatos/leetcode/util"

func PlatesBetweenCandles(s string, queries [][]int) []int {
	m := [][]int{{0, 0}}
	start := -1
	count := 0
	for i := range s {
		if s[i] == '|' {
			if start < 0 {
				count = 0
			}
			m = append(m, []int{i, count})
			start = i
		} else {
			count++
		}
	}

	res := make([]int, 0)
	for i := range queries {
		left := upperBound(m, queries[i][0])
		right := lowerBound(m, queries[i][1])

		res = append(res, util.Max(0, m[right][1]-m[left][1]))
	}

	return res
}

func upperBound(m [][]int, target int) int {
	l, r := 0, len(m)-1
	for l <= r {
		mid := l + (r-l)/2
		if m[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r + 1
}

func lowerBound(m [][]int, target int) int {
	l, r := 0, len(m)-1
	for l <= r {
		mid := l + (r-l)/2
		if m[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l - 1
}

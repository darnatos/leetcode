package solution

import "leetcode/util"

func MaxArea(height []int) int {
	m := 0
	l, r := 0, len(height)-1

	for l < r {
		m = util.Max(m, (r-l)*util.Min(height[l], height[r]))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return m
}

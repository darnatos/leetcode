package solution

import "github.com/darnatos/leetcode/util"

func Candy(ratings []int) int {
	n := len(ratings)
	r2l := make([]int, n)
	for i := range r2l {
		r2l[i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r2l[i] = r2l[i+1] + 1
		}
	}
	pre, cur, res := -1, 0, 0
	for i := range r2l {
		if ratings[i] > pre {
			cur++
		} else {
			cur = 1
		}
		pre = ratings[i]

		res += util.Max(cur, r2l[i])
	}

	return res
}

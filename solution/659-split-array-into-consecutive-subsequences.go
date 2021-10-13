package solution

import (
	"leetcode/util"
	"math"
)

func IsPossible(nums []int) bool {
	p1, p2, p3 := 0, 0, 0
	c1, c2, c3 := 0, 0, 0
	pre, cur, cnt := math.MinInt32, 0, 0

	for i := 0; i < len(nums); pre, p1, p2, p3 = cur, c1, c2, c3 {
		for cur, cnt = nums[i], 0; i < len(nums) && cur == nums[i]; i++ {
			cnt++
		}

		if cur != pre+1 {
			if p1 > 0 || p2 > 0 {
				return false
			}
			c1, c2, c3 = cnt, 0, 0
		} else {
			if cnt < p1+p2 {
				return false
			}
			c1 = util.Max(0, cnt-(p1+p2+p3))
			c2 = p1
			c3 = p2 + util.Min(p3, cnt-(p1+p2))
		}
	}

	return p1 == 0 && p2 == 0
}

package solution

import (
	"github.com/darnatos/leetcode/util"
	"math"
)

func MinSubArrayLen(target int, nums []int) int {
	l, sum := 0, 0
	res := math.MaxInt32
	n := len(nums)

	for i := 0; i < n; i++ {
		sum += nums[i]

		if sum >= target {
			res = util.Min(res, i-l+2)

			for sum >= target {
				sum -= nums[l]
				l++

				res = util.Min(res, i-l+2)
			}
		}
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

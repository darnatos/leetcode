package solution

import (
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	cur, i := 1, 0
	res := int64(0)
	for k > 0 && i < len(nums) {
		if cur > nums[i] {
			i++
			continue
		} else if cur == nums[i] {
			cur++
			continue
		}
		if nums[i]-cur >= k {
			break
		}
		res += int64(nums[i]-1+cur) * int64(nums[i]-cur) / 2
		k -= nums[i] - cur
		cur = nums[i] + 1

	}
	res += int64((cur + cur + k - 1) * k / 2)
	return res
}

package solution

func numSubarrayProductLessThanK(nums []int, k int) int {
	cur := 1
	l := -1
	res := 0
	for i := range nums {
		cur *= nums[i]
		for l < i && cur >= k {
			l++
			cur /= nums[l]
		}
		res += i - l
	}
	return res
}

package solution

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		pre[i] *= suf
		suf *= nums[i]
	}
	return pre
}

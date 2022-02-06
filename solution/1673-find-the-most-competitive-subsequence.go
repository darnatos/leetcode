package solution

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	stack := make([]int, 0, k)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] && len(stack)+n-i > k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}

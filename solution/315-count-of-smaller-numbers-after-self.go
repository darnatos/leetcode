package solution

import "github.com/darnatos/leetcode/solution/BinaryIndexTree"

func countSmaller(nums []int) []int {
	offset := 10001
	bit := BinaryIndexTree.Constructor(2 * offset)
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] += offset
	}
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = bit.Query(nums[i] - 1)
		bit.Add(nums[i], 1)
	}
	return res
}

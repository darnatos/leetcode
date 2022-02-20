package solution

import "github.com/darnatos/leetcode/solution/BinaryIndexTree"

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	m := make([]int, n)
	for i := range nums1 {
		m[nums1[i]] = i
	}
	p := make([]int, n)
	for i := range nums2 {
		p[m[nums2[i]]] = i
	}
	bit := BinaryIndexTree.Constructor(n)
	var res int64
	var left, right int
	for i := 0; i < n; i++ {
		left = bit.Query(p[i])
		right = n - 1 - p[i] - i + left
		res += int64(left * right)
		bit.Add(p[i], 1)
	}
	return res
}

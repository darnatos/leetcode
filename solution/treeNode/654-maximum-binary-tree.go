package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	stack := make([]*TreeNode, 0)
	var cur *TreeNode
	for i := range nums {
		cur = &TreeNode{Val: nums[i]}
		for len(stack) > 0 && stack[len(stack)-1].Val < nums[i] {
			cur.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			stack[len(stack)-1].Right = cur
		}
		stack = append(stack, cur)
	}
	return stack[0]
}

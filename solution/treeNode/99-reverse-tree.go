package treeNode

import "math"

func recoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0)

	pre, cur := &TreeNode{Val: math.MinInt32}, root
	var first *TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if first == nil {
			if pre.Val > cur.Val {
				first = pre
			}
		} else {
			if first.Val < cur.Val {
				first.Val, pre.Val = pre.Val, first.Val
				return
			}
		}
		pre, cur = cur, cur.Right
	}
	if first.Val > pre.Val {
		first.Val, pre.Val = pre.Val, first.Val
	}
}

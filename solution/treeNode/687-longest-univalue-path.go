package treeNode

import (
	"github.com/darnatos/leetcode/util"
)

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	lupDfs(root, &res)
	return res
}

func lupDfs(node *TreeNode, res *int) int {
	l := 0
	if node.Left != nil {
		l = lupDfs(node.Left, res)
		if node.Left.Val == node.Val {
			l++
		} else {
			l = 0
		}
	}
	r := 0
	if node.Right != nil {
		r = lupDfs(node.Right, res)
		if node.Right.Val == node.Val {
			r++
		} else {
			r = 0
		}
	}
	*res = util.Max(*res, l+r)
	return util.Max(l, r)
}

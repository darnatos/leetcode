package treeNode

import (
	"github.com/darnatos/leetcode/util"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return bbtDfs(root) != -1
}

func bbtDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := bbtDfs(root.Left)
	if lh == -1 {
		return -1
	}
	rh := bbtDfs(root.Right)
	if rh == -1 {
		return -1
	}
	if util.Abs(lh-rh) > 1 {
		return -1
	}
	return util.Max(lh, rh) + 1
}

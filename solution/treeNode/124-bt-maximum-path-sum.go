package treeNode

import (
	"github.com/darnatos/leetcode/util"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	mpsHelper(root, &res)
	return res
}

func mpsHelper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := util.Max(mpsHelper(root.Left, res), 0)
	r := util.Max(mpsHelper(root.Right, res), 0)
	m := root.Val

	*res = util.Max(*res, l+m+r)
	return util.Max(l, r) + m
}

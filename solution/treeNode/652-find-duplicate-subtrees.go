package treeNode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []*TreeNode
var fdsMap map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	fdsMap = make(map[string]int)
	res = make([]*TreeNode, 0, 2)
	_ = traverse(root)

	return res
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	s := ""
	s += strconv.Itoa(root.Val) + ","
	s += traverse(root.Left)
	s += traverse(root.Right)
	if _, ok := fdsMap[s]; ok {
		fdsMap[s]++
	} else {
		fdsMap[s] = 1
	}
	if fdsMap[s] == 2 {
		res = append(res, root)
	}
	return s
}

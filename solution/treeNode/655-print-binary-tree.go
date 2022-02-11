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
func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	//1, 3, 7 ...
	w := 1<<h - 1
	res := make([][]string, h)
	for i := range res {
		res[i] = make([]string, w)
	}

	preorder(root, 0, 0, w-1, res)
	return res
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func preorder(root *TreeNode, h, l, r int, res [][]string) {
	if root == nil {
		return
	}
	m := (l + r) / 2
	res[h][m] = strconv.Itoa(root.Val)
	preorder(root.Left, h+1, l, m, res)
	preorder(root.Right, h+1, m+1, r, res)
}

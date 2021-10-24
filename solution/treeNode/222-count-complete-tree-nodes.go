package treeNode

func countNodes(root *TreeNode) int {
	res, l, r := 0, findLeftestRank(root), 0
	for root != nil {
		l, r = l-1, findLeftestRank(root.Right)
		res += 1 << r
		if l == r {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return res
}

func findLeftestRank(root *TreeNode) int {
	res := 0
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

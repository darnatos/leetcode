package treeNode

func isCousins(root *TreeNode, x int, y int) bool {
	xp, xd := icDfs(root, 0, 0, x)
	yp, yd := icDfs(root, 0, 0, y)
	return xp != yp && xd == yd
}

func icDfs(root *TreeNode, pv, d, target int) (int, int) {
	if root == nil {
		return 0, d
	}
	if root.Val == target {
		return pv, d
	}
	d++
	pv = root.Val
	l, ld := icDfs(root.Left, pv, d, target)
	if l == 0 {
		return icDfs(root.Right, pv, d, target)
	}
	return l, ld
}

package treeNode

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ret, _ := ldlDfs(root)
	return ret
}

func ldlDfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return root, 0
	}
	ln, ld := ldlDfs(root.Left)
	rn, rd := ldlDfs(root.Right)
	if ld > rd {
		return ln, ld + 1
	} else if rd > ld {
		return rn, rd + 1
	}
	return root, ld + 1
}

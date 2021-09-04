package treeNode


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if dfs(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func dfs(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot != nil && root.Val == subRoot.Val {
		return dfs(root.Left, subRoot.Left) && dfs(root.Right, subRoot.Right)
	}
	return false
}

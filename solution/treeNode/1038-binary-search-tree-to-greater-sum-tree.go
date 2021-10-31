package treeNode

func bstToGst(root *TreeNode) *TreeNode {
	revDfs(root, 0)
	return root
}

func revDfs(root *TreeNode, cur int) int {
	if root == nil {
		return cur
	}
	root.Val += revDfs(root.Right, cur)
	return revDfs(root.Left, root.Val)
}

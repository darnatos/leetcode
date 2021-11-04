package treeNode

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	sllDfs(root, false, &res)
	return res
}

func sllDfs(root *TreeNode, isLeft bool, res *int) {
	if root == nil {
		return
	}
	if isLeft && root.Left == nil && root.Right == nil {
		*res += root.Val
		return
	}
	sllDfs(root.Left, true, res)
	sllDfs(root.Right, false, res)
}

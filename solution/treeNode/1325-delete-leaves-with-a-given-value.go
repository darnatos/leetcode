package treeNode

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == root.Right && root.Val == target {
		return nil
	}
	return root
}

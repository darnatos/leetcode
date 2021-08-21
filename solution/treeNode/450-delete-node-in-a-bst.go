package treeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {
			root.Val = findMinFromRight(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func findMinFromRight(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

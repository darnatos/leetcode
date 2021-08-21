package treeNode

func Flatten(root *TreeNode) {
	flatten(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	child := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = child
}

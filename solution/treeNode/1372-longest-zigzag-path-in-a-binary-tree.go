package treeNode

func longestZigZag(root *TreeNode) int {
	res := 0
	lzzDfs(root, true, 0, &res)
	lzzDfs(root, false, 0, &res)
	return res
}

func lzzDfs(root *TreeNode, isLeft bool, cur int, res *int) {
	if root == nil {
		return
	}

	if *res < cur {
		*res = cur
	}

	if isLeft {
		lzzDfs(root.Right, false, cur+1, res)
		lzzDfs(root.Left, true, 1, res)
	} else {
		lzzDfs(root.Right, false, 1, res)
		lzzDfs(root.Left, true, cur+1, res)
	}

}

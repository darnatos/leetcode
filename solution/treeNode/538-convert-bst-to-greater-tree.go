package treeNode

func convertBST(root *TreeNode) *TreeNode {
	cur, sum := root, 0
	stack := make([]*TreeNode, 0, 4)

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Val += sum
		sum = cur.Val
		cur = cur.Left
	}

	return root
}

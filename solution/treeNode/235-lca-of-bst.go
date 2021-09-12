package treeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	for {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}

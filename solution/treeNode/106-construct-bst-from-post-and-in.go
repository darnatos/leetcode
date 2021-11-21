package treeNode

func buildTreeInPost(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i := range inorder {
		m[inorder[i]] = i
	}
	rootIdx := len(postorder) - 1
	var recur func(l, r int) *TreeNode
	recur = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		curr := postorder[rootIdx]
		rootIdx--
		node := new(TreeNode)
		node.Val = curr
		mid := m[curr]
		node.Right = recur(mid+1, r)
		node.Left = recur(l, mid-1)
		return node
	}
	return recur(0, len(inorder)-1)
}

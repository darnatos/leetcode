package treeNode

func deepestLeavesSum(root *TreeNode) int {
	var res int
	q := []*TreeNode{root}
	for len(q) != 0 {
		res = 0
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[i]
			res += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
	}

	return res
}

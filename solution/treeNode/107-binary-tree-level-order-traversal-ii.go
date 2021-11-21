package treeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := -1
	cur := 0
	for len(q) != cur {
		level++
		res = append(res, []int{})
		k := len(q)
		for ; cur < k; cur++ {
			res[level] = append(res[level], q[cur].Val)
			if q[cur].Left != nil {
				q = append(q, q[cur].Left)
			}
			if q[cur].Right != nil {
				q = append(q, q[cur].Right)
			}
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}

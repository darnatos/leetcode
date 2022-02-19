package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	m := make([]bool, 1001)
	for _, i := range to_delete {
		m[i] = true
	}
	m[0] = true
	dummy := &TreeNode{Right: root}
	q := make([]*TreeNode, 0)
	q = append(q, dummy)
	res := make([]*TreeNode, 0)
	for len(q) > 0 {
		end := len(q)
		for i := 0; i < end; i++ {
			u := q[i]
			if u.Left != nil {
				q = append(q, u.Left)
				if m[u.Left.Val] {
					u.Left = nil
				} else {
					if m[u.Val] {
						res = append(res, u.Left)
					}
				}
			}
			if u.Right != nil {
				q = append(q, u.Right)
				if m[u.Right.Val] {
					u.Right = nil
				} else {
					if m[u.Val] {
						res = append(res, u.Right)
					}
				}
			}
		}
		q = q[end:]
	}
	return res
}

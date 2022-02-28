package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	q1 := make([]int, 0)
	q = append(q, root)
	q1 = append(q1, 1)
	res := 0
	for len(q) > 0 {
		end := len(q)
		if res < (q1[end-1]-q1[0])+1 {
			res = q1[end-1] - q1[0] + 1
		}
		for i := 0; i < end; i++ {
			cur := q[i]
			cur1 := q1[i]
			if cur != nil {
				if cur.Left != nil {
					q = append(q, cur.Left)
					q1 = append(q1, cur1<<1-1)
				}
				if cur.Right != nil {
					q = append(q, cur.Right)
					q1 = append(q1, cur1<<1)
				}
			}
		}
		q = q[end:]
		q1 = q1[end:]
	}
	return res
}

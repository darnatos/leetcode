package treeNode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	m := make(map[int]int, 8)

	q := make([]*TreeNode, 0, 8)
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		m[node.Val] += 1
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	max := -1
	res := make([]int, 0, 8)
	for k := range m {
		if m[k] > max {
			max = m[k]
			res = append(res[:0], k)
		} else if m[k] == max {
			res = append(res, k)
		}
	}

	return res
}

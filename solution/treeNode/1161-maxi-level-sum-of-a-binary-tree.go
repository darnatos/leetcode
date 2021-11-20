package treeNode

import "math"

func maxLevelSum(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	res := 0
	maxi := math.MinInt32
	level := 0
	i := 0
	for len(q) > i {
		level++
		k := len(q)
		sum := 0
		for ; i < k; i++ {
			sum += q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		if maxi < sum {
			maxi = sum
			res = level
		}
	}
	return res
}

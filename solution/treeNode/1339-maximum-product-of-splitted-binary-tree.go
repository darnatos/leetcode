package treeNode

func maxProduct(root *TreeNode) int {
	sum := calSum(root)
	cur := root
	for {
		if cur.Left != nil {
			if cur.Left.Val > sum-cur.Left.Val {
				cur = cur.Left
				continue
			}
		}
		if cur.Right != nil {
			if cur.Right.Val > sum-cur.Right.Val {
				cur = cur.Right
				continue
			}
		}
		break
	}

	maxMin := sum - cur.Val
	if cur.Left != nil && cur.Left.Val > maxMin {
		maxMin = cur.Left.Val
	}
	if cur.Right != nil && cur.Right.Val > maxMin {
		maxMin = cur.Right.Val
	}
	return maxMin * (sum - maxMin) % (7 + 1e9)
}

func calSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.Val += calSum(root.Left) + calSum(root.Right)

	return root.Val
}

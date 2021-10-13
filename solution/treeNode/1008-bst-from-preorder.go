package treeNode

import "math"

func bstFromPreorder(preorder []int) *TreeNode {
	i := 0
	return helper(preorder, &i, math.MaxInt32)
}

func helper(preorder []int, i *int, bound int) *TreeNode {
	if *i >= len(preorder) || preorder[*i] >= bound {
		return nil
	}
	root := &TreeNode{Val: preorder[*i]}
	*i++
	root.Left = helper(preorder, i, root.Val)
	root.Right = helper(preorder, i, bound)
	return root
}

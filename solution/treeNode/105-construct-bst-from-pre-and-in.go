package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, inorder, 0, 0, len(inorder)-1)
}

func buildTreeHelper(preorder, inorder []int, ps, is, ie int) *TreeNode{
	if is > ie || ps >= len(preorder) {
		return nil
	}
	// fmt.Println(preorder[ps:ps+ie-is], inorder[is:ie])
	root := &TreeNode{preorder[ps], nil, nil}

	ii := 0
	for i := is; i <= ie; i++ {
		if inorder[i] == preorder[ps] {
			ii = i
			break
		}
	}
	root.Left = buildTreeHelper(preorder, inorder, ps+1, is, ii-1)
	root.Right = buildTreeHelper(preorder, inorder, ps+ii-is+1, ii+1, ie)
	return root
}

package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
	sp := make([]byte, 0, 4)
	gdHelper(root, startValue, &sp)
	dp := make([]byte, 0, 4)
	gdHelper(root, destValue, &dp)

	reverseBytes(sp)
	reverseBytes(dp)

	i := 0
	for ; i < len(sp) && i < len(dp); i++ {
		if dp[i] != sp[i] {
			break
		}
	}
	for j := range sp {
		sp[j] = 'U'
	}
	return string(append(sp[i:], dp[i:]...))
}

func gdHelper(root *TreeNode, target int, path *[]byte) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}
	l := gdHelper(root.Left, target, path)
	if l {
		*path = append(*path, 'L')
		return true
	}
	r := gdHelper(root.Right, target, path)
	if r {
		*path = append(*path, 'R')
	}
	return r
}

func reverseBytes(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

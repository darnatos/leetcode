package treeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int][]int)
	r := make(map[int]bool)
	for _, desc := range descriptions {
		if _, ok := m[desc[0]]; !ok {
			m[desc[0]] = make([]int, 2)
		}
		if _, ok := r[desc[0]]; !ok {
			r[desc[0]] = true
		}
		r[desc[1]] = false
		m[desc[0]][1-desc[2]] = desc[1]
	}
	rv := 0
	for i := range r {
		if r[i] == true {
			rv = i
			break
		}
	}

	return dfsCbt(m, rv)
}

func dfsCbt(m map[int][]int, cur int) *TreeNode {
	if cur == 0 {
		return nil
	}
	var l, r *TreeNode
	if len(m[cur]) > 0 {
		l = dfsCbt(m, m[cur][0])
		r = dfsCbt(m, m[cur][1])
	}
	return &TreeNode{
		Val:   cur,
		Left:  l,
		Right: r,
	}
}

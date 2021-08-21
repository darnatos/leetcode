package treeNode

const MaxValue = 2147483648

func findSecondMinimumValue(root *TreeNode) int {
	hSet := make(map[int]bool, 32)
	Traverse(root, hSet)

	min := MaxValue

	for i := range hSet {
		if i != root.Val {
			min = Min(min, i)
		}
	}

	if min == MaxValue {
		return -1
	}
	return min
}

func Traverse(root *TreeNode, hSet map[int]bool) {
	hSet[root.Val] = true
	if root.Left != nil {
		Traverse(root.Left, hSet)
		Traverse(root.Right, hSet)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

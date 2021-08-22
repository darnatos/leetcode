package treeNode

import "sort"

type Pair struct {
	row int
	val int
}

type PairList []Pair

func (arr PairList) Len() int {
	return len(arr)
}

func (arr PairList) Less(i, j int) bool {
	if arr[i].row < arr[j].row {
		return true
	} else if arr[i].row == arr[j].row && arr[i].val < arr[j].val {
		return true
	}
	return false
}

func (arr PairList) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr PairList) GetVal(i int) int {
	return arr[i].val
}

var m map[int]PairList

func verticalTraversal(root *TreeNode) [][]int {
	ans := make([][]int, 0, 16)
	if root == nil {
		return ans
	}

	m = make(map[int]PairList)

	populate(root, 0, 0)

	min := len(m)
	for i := range m {
		if min > i {
			min = i
		}
	}

	for i := min; i < len(m)+min; i++ {
		sort.Sort(m[i])
		res := make([]int, 0, len(m[i]))
		for j := range m[i] {
			res = append(res, m[i].GetVal(j))
		}
		ans = append(ans, res)
	}
	return ans
}

func populate(node *TreeNode, row, col int) {
	if node == nil {
		return
	}
	if _, ok := m[col]; !ok {
		m[col] = make(PairList, 0)
	}
	m[col] = append(m[col], Pair{row, node.Val})
	populate(node.Left, row+1, col-1)
	populate(node.Right, row+1, col+1)
}

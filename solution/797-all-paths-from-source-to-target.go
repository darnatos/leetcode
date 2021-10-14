package solution

func AllPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	apstDfs(graph, 0, []int{}, &res)
	return res
}

func apstDfs(graph [][]int, i int, cur []int, res *[][]int) {
	cur = append(cur, i)
	if i == len(graph)-1 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
	} else {
		for _, j := range graph[i] {
			apstDfs(graph, j, cur, res)
		}
	}
	cur = cur[:len(cur)-1]
}

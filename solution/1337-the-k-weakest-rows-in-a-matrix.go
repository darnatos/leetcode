package solution

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	arr := make([][]int, len(mat))
	for i := range mat {
		arr[i] = []int{0, i}
		for j := range mat[i] {
			arr[i][0] += mat[i][j]
		}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] || arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1] })
	res := make([]int, 0, k)
	for i := range arr {
		res = append(res, arr[i][1])
		if len(res) == k {
			break
		}
	}
	return res
}

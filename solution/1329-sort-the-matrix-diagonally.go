package solution

import "sort"

func DiagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	arr := make([]int, 0, m)
	for i := range mat {
		arr = arr[:0]
		for j, k := i, 0; j < m && k < n; j, k = j+1, k+1 {
			arr = append(arr, mat[j][k])
		}
		sort.Ints(arr)
		for j, k := i, 0; j < m && k < n; j, k = j+1, k+1 {
			mat[j][k] = arr[k]
		}
	}

	for i := 1; i < n; i++ {
		arr = arr[:0]
		for j, k := 0, i; j < m && k < n; j, k = j+1, k+1 {
			arr = append(arr, mat[j][k])
		}
		sort.Ints(arr)
		for j, k := 0, i; j < m && k < n; j, k = j+1, k+1 {
			mat[j][k] = arr[j]
		}
	}
	return mat
}

package solution

func Construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = original[i*n : (i+1)*n]
	}
	return res
}

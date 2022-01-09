package solution

func checkValid(matrix [][]int) bool {
	for i := range matrix {
		existsR := make([]bool, len(matrix)+1)
		existsC := make([]bool, len(matrix)+1)
		for j := range matrix {
			if existsR[matrix[i][j]] || existsC[matrix[j][i]] {
				return false
			}
			existsR[matrix[i][j]] = true
			existsC[matrix[j][i]] = true
		}
	}
	return true
}

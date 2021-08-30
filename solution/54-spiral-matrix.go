package solution

func SpiralOrder(matrix [][]int) []int {
	e := [4]int{1, len(matrix[0]) - 1, len(matrix) - 1, 0}
	flag := 0
	res := make([]int, len(matrix)*len(matrix[0]))

	i, j := 0, -1
	for k := range res {
		switch flag {
		case 0:
			if j < e[1] {
				j++
			} else if j == e[1] {
				i++
				e[1]--
				flag = 1
			}
		case 1:
			if i < e[2] {
				i++
			} else if i == e[2] {
				j--
				e[2]--
				flag = 2
			}
		case 2:
			if j > e[3] {
				j--
			} else if j == e[3] {
				i--
				e[3]++
				flag = 3
			}
		case 3:
			if i > e[0] {
				i--
			} else if i == e[0] {
				j++
				e[0]++
				flag = 0
			}
		}
		res[k] = matrix[i][j]
	}

	return res
}

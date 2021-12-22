package solution

func numSubmat(mat [][]int) int {
	h := make([]int, len(mat[0]))
	res := 0
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				h[j] += 1
			} else {
				h[j] = 0
			}
		}
		sum := make([]int, len(h))
		stack := make([]int, 0, len(h))
		for j := range h {
			for len(stack) > 0 && h[stack[len(stack)-1]] >= h[j] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				k := stack[len(stack)-1]
				sum[j] = sum[k]
				sum[j] += h[j] * (j - k)
			} else {
				sum[j] = h[j] * (j + 1)
			}
			stack = append(stack, j)
		}
		for i := range sum {
			res += sum[i]
		}
	}
	return res
}

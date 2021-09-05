package solution

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	res := 1
	for i, j := m, 1; j < n; i, j = i+1, j+1 {
		res *= i
		res /= j
	}
	return res
}

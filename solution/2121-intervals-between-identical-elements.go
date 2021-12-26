package solution

func getDistances(arr []int) []int64 {
	m := make(map[int][]int, len(arr))
	for i := range arr {
		m[arr[i]] = append(m[arr[i]], i)
	}

	res := make([]int64, len(arr))

	for i := range m {
		sum := int64(0)
		m0 := m[i][0]
		for _, j := range m[i] {
			sum += int64(j - m0)
		}

		res[m0] = sum
		for j := 1; j < len(m[i]); j++ {
			sum += int64((j + j - len(m[i])) * (m[i][j] - m[i][j-1]))
			res[m[i][j]] = sum
		}
	}
	return res
}

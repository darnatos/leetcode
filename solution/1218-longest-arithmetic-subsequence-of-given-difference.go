package solution

func LongestSubsequence(arr []int, difference int) (res int) {
	m := make(map[int]int)

	for _, v := range arr {
		if val, ok := m[v-difference]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
		if m[v] > res {
			res = m[v]
		}
	}
	return
}

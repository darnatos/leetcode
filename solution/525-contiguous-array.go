package solution

func findMaxLength(nums []int) int {
	n := len(nums)
	res := 0
	cnt := 0
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < n; i++ {
		cnt += nums[i]*2 - 1
		if _, ok := m[cnt]; ok {
			res = max(res, i-m[cnt])
		} else {
			m[cnt] = i
		}
	}
	return res
}

package solution

func findPairs(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 0
			if k == 0 {
				continue
			}
		}
		if v, ok := m[nums[i]-k]; ok && v&1 == 0 {
			m[nums[i]] |= 2
			m[nums[i]-k] |= 1
			res++
		}
		if v, ok := m[nums[i]+k]; ok && v&2 == 0 {
			m[nums[i]] |= 1
			m[nums[i]+k] |= 2
			res++
		}
	}
	return res
}

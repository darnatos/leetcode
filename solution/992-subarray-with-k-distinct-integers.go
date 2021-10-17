package solution

func SubarraysWithKDistinct(nums []int, k int) int {
	m := make([]int, len(nums)+1)
	res := 0
	for i, j, prefix := 0, 0, 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			k--
		}
		m[nums[i]]++
		if k < 0 {
			m[nums[j]]--
			j++
			k++
			prefix = 0
		}
		for m[nums[j]] > 1 {
			m[nums[j]]--
			j++
			prefix++
		}
		if k == 0 {
			res += prefix + 1
		}
	}

	return res
}

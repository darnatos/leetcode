package solution

func NumberOfSubarrays(nums []int, k int) int {
	return fo(nums, k) - fo(nums, k-1)
}

func fo(nums []int, a int) int {
	s := 0
	i, j := 0, 0
	ans := 0
	for j < len(nums) {
		if nums[j]&1 == 1 {
			s++
		}

		for s > a {
			if nums[i]&1 == 1 {
				s--
			}
			i++
		}
		ans += j - i + 1
		j++
	}
	return ans
}
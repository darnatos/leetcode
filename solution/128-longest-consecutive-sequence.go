package solution

func LongestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for i := range nums {
		m[nums[i]] = true
	}

	res := 0
	for i := range m {
		if m[i-1] {
			continue
		}
		j := i+1
		for m[j] {
			j++
		}
		if res < j-i {
			res = j-i
		}
	}

	return res
}

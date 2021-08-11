package solution

func RemoveDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	k := 1
	for i := 1; i < size; i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k = k + 1
		}
	}

	return k
}

package solution

func NextPermutation(nums []int) {
	k := npFindK(nums)

	if k == -1 {
		reverseInts(nums)
		return
	}

	for l := len(nums) - 1; l > k; l-- {
		if nums[k] < nums[l] {
			nums[k], nums[l] = nums[l], nums[k]
			break
		}
	}
	reverseInts(nums[k+1:])
}

func npFindK(nums []int) int {
	for k := len(nums) - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			return k
		}
	}
	return -1
}

func reverseInts(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

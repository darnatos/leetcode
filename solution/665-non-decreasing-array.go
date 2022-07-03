package solution

func checkPossibility(nums []int) bool {
	count := 0
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			index = i
			count++
		}
	}
	if count == 0 {
		return true
	}
	if count == 1 {
		if index == 1 || index == len(nums)-1 {
			return true
		}
		if nums[index+1] >= nums[index-1] || nums[index] >= nums[index-2] {
			return true
		}
	}
	return false
}

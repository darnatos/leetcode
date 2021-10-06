package solution

func FindDuplicates(nums []int) []int {
	nums = append(nums, 0)
	ret := make([]int, 0, len(nums)/2)
	for j := range nums {
		for nums[j] != 0 && j != nums[j] {
			if nums[j] == nums[nums[j]] {
				ret = append(ret, nums[j])
				nums[j], nums[nums[j]] = 0, 0
			} else {
				nums[j], nums[nums[j]] = nums[nums[j]], nums[j]
			}
		}
	}
	return ret
}

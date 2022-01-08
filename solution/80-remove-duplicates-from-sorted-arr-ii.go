package solution

func removeDuplicates(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if res < 2 || nums[i] != nums[res-2] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

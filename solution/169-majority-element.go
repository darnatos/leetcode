package solution

func MajorityElement(nums []int) int {
	c := len(nums)/2
	major, count := nums[0], 1
	for i := 1; i < len(nums) && count <= c; i++ {
		if count == 0 {
			count++
			major = nums[i]
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

package solution

func FindMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[r] > nums[m] {
			r = m
		} else if nums[r] < nums[m] {
			l = m + 1
		} else {
			r--
		}
	}
	return nums[l]
}

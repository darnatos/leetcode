package solution

func FindPeakElement(nums []int) int {
	return fpeHelper(nums, 0, len(nums)-1)
}

func fpeHelper(nums []int, l, r int) int {
	if l == r {
		return l
	}
	m := (l+r)/2

	if nums[m] > nums[m+1] {
		return fpeHelper(nums, l, m)
	} else {
		return fpeHelper(nums, m+1, r)
	}
}


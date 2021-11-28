package solution

func MinimumDeletions(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	mni, mxi := 0, 0
	for i := range nums {
		if nums[i] < nums[mni] {
			mni = i
		}
		if nums[i] > nums[mxi] {
			mxi = i
		}
	}
	res := 0
	a, b := mni, mxi
	if a > b {
		a, b = b, a
	}
	res = min(min(b+1, a+1+len(nums)-b), len(nums)-a)
	return res
}

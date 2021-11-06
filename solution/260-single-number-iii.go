package solution

func SingleNumber3(nums []int) []int {
	aXORb := 0
	for i := range nums {
		aXORb = aXORb ^ nums[i]
	}

	mask := aXORb & (-aXORb)
	res := []int{0, 0}
	for i := range nums {
		if mask&nums[i] != 0 {
			res[0] ^= nums[i]
		} else {
			res[1] ^= nums[i]
		}
	}

	return res
}

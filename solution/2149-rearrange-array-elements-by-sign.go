package solution

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	e, o := 0, 1
	for i := range nums {
		if nums[i] > 0 {
			res[e] = nums[i]
			e += 2
		} else {
			res[o] = nums[i]
			o += 2
		}
	}
	return res
}

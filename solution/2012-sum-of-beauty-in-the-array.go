package solution

import "math"

func SumOfBeauties(nums []int) int {
	mins := make(map[int]int, len(nums))
	tm := math.MaxInt32
	mins[len(nums)-1] = tm
	for i := len(nums)-2; i >= 0 ; i-- {
		if tm > nums[i+1] {
			tm = nums[i+1]
		}
		mins[i] = tm
	}

	res, m := 0, nums[0]
	for i := 1; i <= len(nums)-2; i++ {
		if nums[i] > m && nums[i] < mins[i] {
			res += 2
			m = nums[i]
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res += 1
		}
		m = max(m, nums[i])
	}
	return res
}

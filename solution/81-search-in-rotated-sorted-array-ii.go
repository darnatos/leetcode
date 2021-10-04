package solution

func SearchII(nums []int, target int) bool {
	//nums[r] <= nums[l] || nums[m] <= nums[r]
	l, r, m := 0, len(nums)-1, 0
	for l <= r {
		m = (l + r) >> 1
		if target == nums[m] {
			return true
		}

		if nums[l] == nums[r] && nums[l] == nums[m] {
			l++
			r--
		} else if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return false
}


package solution

import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	cd := abs(res - target)
	for i := range nums {
		k := target - nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if l == i {
				l++
				continue
			}
			if r == i {
				r--
				continue
			}
			tmp := nums[l] + nums[r]
			sum := nums[i] + tmp
			if tmp > k {
				r--
				tmp -= k
			} else {
				l++
				tmp = k - tmp
			}
			if tmp < cd {
				cd = tmp
				res = sum
			}
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

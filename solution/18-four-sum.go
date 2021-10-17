package solution

import "sort"

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4, 0)
}

func kSum(nums []int, target, k, l int) [][]int {
	n := len(nums)
	res := make([][]int, 0)

	if k == 2 {
		r := n - 1
		for l < r {
			if target == nums[l]+nums[r] {
				tmp := make([]int, 0, 4)
				tmp = append(tmp, nums[r], nums[l])
				res = append(res, tmp)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if target > nums[l]+nums[r] {
				l++
			} else {
				r--
			}
		}
	} else {
		for i := l; i < n-k+1; i++ {
			tmp := kSum(nums, target-nums[i], k-1, i+1)
			if len(tmp) != 0 {
				for j := range tmp {
					tmp[j] = append(tmp[j], nums[i])
				}
				res = append(res, tmp...)
			}
			for i < n-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

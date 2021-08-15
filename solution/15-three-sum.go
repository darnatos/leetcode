package solution

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return res
		}
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if 0 == sum {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for k > j && nums[k+1] == nums[k] {
					k--
				}
				for k > j && nums[j-1] == nums[k] {
					j++
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

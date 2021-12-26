package solution

import "sort"

func recoverArray(nums []int) []int {
	n := len(nums) / 2
	res := make([]int, 0, n)
	sort.Ints(nums)

	for i := 1; i < 2*n; i++ {
		k := nums[i] - nums[0]
		if k == 0 || k%2 == 1 {
			continue
		}
		res = res[:0]
		count := make(map[int]int, 2*n)
		for i := range nums {
			count[nums[i]]++
		}

		for j := range nums {
			if count[nums[j]] > 0 {
				count[nums[j]]--
				if count[nums[j]+k] > 0 {
					count[nums[j]+k]--
					res = append(res, nums[j]+k/2)
					if len(res) >= n {
						return res
					}
				} else {
					break
				}
			}
		}
	}

	return res
}

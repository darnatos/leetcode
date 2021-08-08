package solution

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		if i != m[v] && nums[i]+nums[m[v]] == target {
			return []int{m[v], i}
		}
		m[v] = i
	}
	sort.Ints(nums)
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else if sum == target {
			r1, r2 := m[nums[lo]], m[nums[hi]]
			if r1 > r2 {
				return []int{r2, r1}
			}
			return []int{r1, r2}
		}
	}
	return nil
}

package solution

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	snums := make([]string, len(nums))
	for i := range nums {
		snums[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(snums, func(a, b int) bool {
		s1 := snums[a]+snums[b]
		s2 := snums[b]+snums[a]
		return s1 > s2
	})

	res := ""
	for i := range snums {
		res += snums[i]
	}
	i := 0
	for ; i < len(res)-1; i++ {
		if res[i] != '0' {
			break
		}
	}
	res = res[i:]

	return res
}

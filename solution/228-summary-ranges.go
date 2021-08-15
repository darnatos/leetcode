package solution

import "fmt"

func SummaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}

	s, e := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			e = i - 1
			if s == e {
				res = append(res, fmt.Sprintf("%d", nums[s]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[s], nums[e]))
			}
			s, e = i, i
		}
	}
	e = len(nums) - 1
	if s == e {
		res = append(res, fmt.Sprintf("%d", nums[s]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[s], nums[e]))
	}

	return res
}

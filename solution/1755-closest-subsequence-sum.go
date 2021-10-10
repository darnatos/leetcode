package solution

import (
	"math"
	"sort"
)

func MinAbsDifference(nums []int, goal int) int {
	n := len(nums)
	first, second := make([]int, 0, n/2+1), make([]int, 0, n/2+1)
	generateABD(nums, 0, n/2, 0, &first)
	generateABD(nums, n/2, n, 0, &second)
	sort.Ints(first)
	ret := math.MaxInt32

	for _, secondSetSum := range second {
		left := goal - secondSetSum
		if first[0] > left {
			ret = min(ret, abs(first[0]+secondSetSum-goal))
			continue
		}
		if first[len(first)-1] < left {
			ret = min(ret, abs(first[len(first)-1]+secondSetSum-goal))
			continue
		}
		pos := binarySearch(first, left)
		if pos >= 0 {
			return 0
		} else {
			pos = -(pos + 1)
		}
		if pos > 0 {
			low := pos - 1
			ret = min(ret, abs(first[low]+secondSetSum-goal))
		}
		ret = min(ret, abs(first[pos]+secondSetSum-goal))
	}
	return ret
}

func generateABD(nums []int, i, end, sum int, list *[]int) {
	if i == end {
		*list = append(*list, sum)
		return
	}
	generateABD(nums, i+1, end, sum, list)
	generateABD(nums, i+1, end, sum+nums[i], list)
}

func binarySearch(nums []int, key int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		if nums[m] == key {
			return m
		} else if nums[m] > key {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -l-1
}

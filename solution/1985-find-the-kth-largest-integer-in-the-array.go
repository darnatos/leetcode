package solution

import (
	"math/rand"
	"time"
)

func KthLargestNumber(nums []string, k int) string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	k--
	l, r := 0, len(nums)-1
	for l <= r {
		m := klnPartition(nums, l, r)
		if m == k {
			break
		}
		if m > k {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return nums[k]
}

func klnPartition(nums []string, l, r int) int {
	pivot := nums[l]
	for l < r {
		for l < r && !compare(nums[r], pivot) {
			r--
		}
		nums[l] = nums[r]
		for l < r && compare(nums[l], pivot) {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}

func compare(a, b string) bool {
	if len(a) == len(b) {
		return a > b
	}
	return len(a) > len(b)
}

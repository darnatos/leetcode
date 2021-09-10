package solution

import "math/rand"

var s1 rand.Source
var r1 *rand.Rand

func SortArray(nums []int) []int {
	s1 = rand.NewSource(1)
	r1 = rand.New(s1)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l < r {
		k := l + r1.Intn(r-l)
		nums[k], nums[r] = nums[r], nums[k]

		pi := partition(nums, l, r)
		quickSort(nums, l, pi-1)
		quickSort(nums, pi+1, r)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

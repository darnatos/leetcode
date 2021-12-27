package solution

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	k -= 1
	stack := make([]int, 0)

	stack = append(stack, 0, len(nums)-1)

	rand.Seed(time.Now().UnixNano())

	for len(stack) != 0 {
		l, r := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if l < r {

			pivotIndex := rand.Intn(r-l+1) + l
			nums[r], nums[pivotIndex] = nums[pivotIndex], nums[r]

			pivot := nums[r]
			i := l
			for j := l; j <= r; j++ {
				if nums[j] > pivot {
					nums[j], nums[i] = nums[i], nums[j]
					i++
				}
			}
			nums[r], nums[i] = nums[i], nums[r]

			if i > k {
				stack = append(stack, l, i-1)
			} else if i < k {
				stack = append(stack, i+1, r)
			} else {
				break
			}
		}
	}
	return nums[k]
}

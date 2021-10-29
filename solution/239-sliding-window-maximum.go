package solution

func MaxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		for len(queue) != 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}

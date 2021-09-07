package solution

func Jump2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	jumps, l, r := 1, 0, nums[0]
	for r < len(nums)-1 {
		jumps++
		n := l + nums[l]
		for i := l; i <= r; i++ {
			n = max(n, i+nums[i])
		}
		l, r = r, n
	}
	return jumps
}

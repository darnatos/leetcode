package solution

func minOperations(nums []int, x int) int {
	n := len(nums)
	l, r := 0, 0
	sum, res := 0, n+1
	for i := range nums {
		sum += nums[i]
	}
	for l <= r {
		if sum >= x {
			if sum == x {
				res = min(res, n-r+l)
			}
			if r < n {
				sum -= nums[r]
				r++
			} else {
				break
			}
		} else {
			sum += nums[l]
			l++
		}
	}
	if res > n {
		return -1
	}
	return res
}

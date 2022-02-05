package solution

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		if nums[i] < pivot {
			res[l] = nums[i]
			l++
		}
		if nums[j] > pivot {
			res[r] = nums[j]
			r--
		}
	}
	for l <= r {
		res[l] = pivot
		l++
	}

	return res
}

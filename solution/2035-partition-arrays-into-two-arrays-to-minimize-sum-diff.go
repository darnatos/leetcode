package solution

import "sort"

func MinimumDifference(nums []int) int {
	n, sum := len(nums)/2, 0
	for i := range nums {
		sum += nums[i]
	}

	left, right := make([][]int, n+1), make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		left[i] = make([]int, 0)
		right[i] = make([]int, 0)
	}
	for i := 0; i < 1<<n; i++ {
		sz, l, r := 0, 0, 0
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				sz++
				l += nums[j]
				r += nums[n+j]
			}
		}
		left[sz] = append(left[sz], l)
		right[sz] = append(right[sz], r)
	}
	for sz := range right {
		sort.Ints(right[sz])
	}

	res := min(abs(sum-2*left[n][0]), abs(sum-2*right[n][0]))
	// fmt.Println(left, right, res)
	for i := 1; i < n; i++ {
		for _, a := range left[i] {
			b, j := (sum-2*a)/2, n-i
			if b > right[j][len(right[j])-1] {
				res = min(res, abs(sum-2*(a+right[j][len(right[j])-1])))
				continue
			}
			if b < right[j][0] {
				res = min(res, abs(sum-2*(a+right[j][0])))
				continue
			}
			pos := binarySearch(right[j], b)
			if pos < 0 {
				pos = -(pos + 1)
			}
			// fmt.Println(i, j, sum, a, pos, right[j])
			if pos > 0 {
				low := pos - 1
				res = min(res, abs(sum-2*(a+right[j][low])))
			}
			res = min(res, abs(sum-2*(a+right[j][pos])))
		}
	}
	return res
}

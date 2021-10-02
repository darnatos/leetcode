package solution

func WaysToPartition(nums []int, k int) int {
	n := len(nums)
	pref, suff := make([]int, n), make([]int, n)
	pref[0], suff[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + nums[i]
		suff[n-1-i] = suff[n-i] + nums[n-1-i]
	}

	res := 0
	left, right := make(map[int]int, n), make(map[int]int, n)

	for i := 0; i < n-1; i++ {
		right[pref[i]-suff[i+1]]++
	}
	if v, ok := right[0]; ok {
		res = v
	}

	for i := 0; i < n; i++ {
		cur, diff := 0, k-nums[i]
		if v, ok := left[diff]; ok {
			cur += v
		}
		if v, ok := right[-diff]; ok {
			cur += v
		}
		if res < cur {
			res = cur
		}
		if i < n-1 {
			dd := pref[i] - suff[i+1]
			right[dd]--
			left[dd]++
			if right[dd] == 0 {
				delete(right, dd)
			}
		}
		// fmt.Println(cur, diff, res, left, right)
	}

	return res
}

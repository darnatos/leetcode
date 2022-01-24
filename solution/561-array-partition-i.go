package solution

func arrayPairSum(nums []int) int {
	bucket := make([]int, 20001)
	for i := range nums {
		bucket[nums[i]+10000]++
	}

	res, even := 0, true
	for i := 0; i < 20001; {
		if bucket[i] > 0 {
			if even {
				res += i - 10000
			}
			even = !even
			bucket[i]--
		} else {
			i++
		}
	}
	return res
}

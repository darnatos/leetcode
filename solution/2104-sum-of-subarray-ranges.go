package solution

func subArrayRanges(nums []int) int64 {
	var sum int64 = 0
	for i := 0; i < len(nums); i++ {
		var l, h int64 = 1000000000, -10000000000
		for j := i; j < len(nums); j++ {
			k := int64(nums[j])
			if l > k {
				l = k
			}
			if h < k {
				h = k
			}
			sum += h - l
		}
	}
	return sum
}

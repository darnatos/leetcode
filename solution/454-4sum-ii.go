package solution

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ab := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			ab[a+b]++
		}
	}
	res := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			res += ab[-c-d]
		}
	}
	return res
}

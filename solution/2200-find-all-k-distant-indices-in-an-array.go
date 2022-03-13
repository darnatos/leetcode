package solution

func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0)
	for i := range nums {
		if nums[i] == key {
			start := max(0, i-k)
			if len(res) > 0 {
				start = max(start, res[len(res)-1]+1)
			}
			for j := start; j <= min(i+k, len(nums)-1); j++ {
				res = append(res, j)
			}
		}
	}
	return res
}

package solution

func GetAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	r := 2*k + 1
	sum := 0
	cnt := 0
	for i := 0; i < r && i < len(nums); i++ {
		cnt++
		sum += nums[i]
	}
	for i := range nums {
		if i < k || i > len(nums)-1-k || cnt < r {
			res[i] = -1
		} else {
			if i > k {
				sum += nums[i+k] - nums[i-k-1]
			}
			res[i] = sum / r
		}
	}
	return res
}

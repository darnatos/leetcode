package solution

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	cnt := make(map[int]int)
	preSum := 0
	cnt[preSum] = -1
	for i := 0; i < len(nums); i++ {
		preSum = (preSum + nums[i]) % k
		if v, ok := cnt[preSum]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			cnt[preSum] = i
		}
	}
	return false
}

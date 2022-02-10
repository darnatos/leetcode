package solution

func subarraysDivByK(nums []int, k int) int {
	pre := make(map[int]int)
	res, sum := 0, 0
	pre[sum]++
	for i := range nums {
		sum = (sum + nums[i]) % k
		if sum < 0 {
			sum += k
		}
		res += pre[sum]
		pre[sum]++
	}
	return res
}

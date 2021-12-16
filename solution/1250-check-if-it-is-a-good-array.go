package solution

func isGoodArray(nums []int) bool {
	res := nums[0]
	for _, n := range nums {
		res = gcd(res, n)
	}
	return res == 1
}

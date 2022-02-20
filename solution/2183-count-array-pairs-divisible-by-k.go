package solution

func coutPairs(nums []int, k int) int64 {
	n := len(nums)
	divisors := make([]int, 0)
	for i := 1; i <= k; i++ {
		if k%i == 0 {
			divisors = append(divisors, i)
		}
	}
	cnt := make([]int, k+1)
	var res int64 = 0
	for i := 0; i < n; i++ {
		remainder := k / gcd(k, nums[i])
		res += int64(cnt[remainder])
		for _, div := range divisors {
			if nums[i]%div == 0 {
				cnt[div]++
			}
		}
	}
	return res
}

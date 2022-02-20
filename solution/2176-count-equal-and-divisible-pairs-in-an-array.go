package solution

func countPairs(nums []int, k int) int {
	res := 0
	cnt := make(map[int]int)
	for i := range nums {
		if i%k == 0 {
			res += cnt[nums[i]]
		} else {
			d := k / gcd(i, k)
			for j := 0; j < i; j += d {
				if nums[i] == nums[j] {
					res++
				}
			}
		}
		cnt[nums[i]]++
	}
	return res
}

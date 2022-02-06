package solution

func countDifferentSubsequenceGCDs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hash := make(map[int]struct{})
	limit := 0
	for i := range nums {
		if limit < nums[i] {
			limit = nums[i]
		}
		hash[nums[i]] = struct{}{}
	}

	res := 0
	for i := 1; i <= limit; i++ {
		g := 0
		for j := i; j <= limit; j += i {
			if _, ok := hash[j]; ok {
				g = gcd(g, j)
			}
		}
		if g == i {
			res++
		}
	}

	return res
}

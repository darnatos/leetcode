package solution

func countQuadruplets(nums []int) int {
	res := 0
	count := make(map[int]int)
	for b := range nums {
		for a := 0; a < b; a++ {
			count[nums[a]+nums[b]]++
		}
		c := b + 1
		for d := c + 1; d < len(nums); d++ {
			res += count[nums[d]-nums[c]]
		}
	}
	return res
}

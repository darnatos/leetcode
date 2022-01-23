package solution

func findLonely(nums []int) []int {
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]]++
	}
	res := make([]int, 0)
	for i := range nums {
		if cnt[nums[i]] == 1 {
			if cnt[nums[i]+1] == 0 && cnt[nums[i]-1] == 0 {
				res = append(res, nums[i])
			}
		}
	}
	return res
}

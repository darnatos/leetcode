package solution

func deleteAndEarn(nums []int) int {
	cnt := make([]int, 10001)
	for i := range nums {
		cnt[nums[i]] += nums[i]
	}

	take, skip := 0, 0
	for i := range cnt {
		take, skip = skip+cnt[i], max(take, skip)
	}
	return max(take, skip)
}

package solution

func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}

	chosen := make([]bool, len(nums))

	return cpksBacktrack(nums, 0, 0, k, sum/k, chosen)
}

func cpksBacktrack(nums []int, start, sum, k, target int, chosen []bool) bool {
	if k == 1 {
		return true
	}
	if sum > target {
		return false
	}
	if sum == target {
		return cpksBacktrack(nums, 0, 0, k-1, target, chosen)
	}

	for i := start; i < len(nums); i++ {
		if chosen[i] {
			continue
		}
		chosen[i] = true
		if cpksBacktrack(nums, i+1, sum+nums[i], k, target, chosen) {
			return true
		}
		chosen[i] = false
	}

	return false
}

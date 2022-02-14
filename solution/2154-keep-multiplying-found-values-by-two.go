package solution

func findFinalValue(nums []int, original int) int {
	exist := [1001]bool{}
	for i := range nums {
		exist[nums[i]] = true
	}
	for original <= 1000 && exist[original] {
		original *= 2
	}
	return original
}

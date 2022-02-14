package solution

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	right := 0
	for i := range nums {
		if nums[i] == 1 {
			right++
		}
	}
	maxScore := right
	res := make([]int, 0, len(nums)/2)
	res = append(res, 0)
	sum := right
	for i := 0; i < n; i++ {
		sum += 1 - nums[i]<<1
		if sum > maxScore {
			maxScore = sum
			res = res[:0]
		}
		if sum == maxScore {
			res = append(res, i+1)
		}
	}
	return res
}

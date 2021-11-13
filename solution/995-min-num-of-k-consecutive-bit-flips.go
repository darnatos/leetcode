package solution

func MinKBitFlips(nums []int, k int) int {
	res := 0
	isFlipped := 0
	for i := 0; i < len(nums); i++ {
		if isFlipped&1 == nums[i] {
			if i+k-1 >= len(nums) {
				return -1
			}
			res++
			isFlipped ^= 1
			nums[i] -= 2
		}
		if i >= k-1 && nums[i-k+1] < 0 {
			isFlipped ^= 1
			nums[i-k+1] += 2
		}
	}
	return res
}

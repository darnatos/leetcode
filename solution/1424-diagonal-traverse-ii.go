package solution

func FindDiagonalOrder(nums [][]int) []int {
	bucket := make([][]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j < len(bucket) {
				bucket[i+j] = append(bucket[i+j], nums[i][j])
			} else {
				bucket = append(bucket, []int{nums[i][j]})
			}
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(bucket); i++ {
		for j := len(bucket[i]) - 1; j >= 0; j-- {
			res = append(res, bucket[i][j])
		}
	}
	return res
}

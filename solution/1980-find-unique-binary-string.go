package solution

func findDifferentBinaryString(nums []string) string {
	res := make([]byte, len(nums[0]))
	for i := range nums {
		if nums[i][i] == '1' {
			res[i] = '0'
		} else {
			res[i] = '1'
		}
	}
	return string(res)
}

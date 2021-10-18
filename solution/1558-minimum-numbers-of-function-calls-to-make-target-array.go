package solution

func MinOperationsFCTA(nums []int) int {
	res, mb := 0, 1
	for _, v := range nums {
		bits := 0
		for v > 0 {
			res += v & 1
			v >>= 1
			bits++
		}
		if mb < bits {
			mb = bits
		}
	}
	return res + mb - 1
}

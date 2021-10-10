package solution

func RangeBitwiseAnd(left int, right int) int {
	if left < right {
		return RangeBitwiseAnd(left>>1, right>>1) << 1
	}
	return right
}

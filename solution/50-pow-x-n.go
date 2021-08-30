package solution

func myPow(x float64, n int) float64 {
	var res float64 = 1
	if n > 0 {
		for i := 0; i < n; i++ {
			res *= x
		}
	} else {
		for i := 0; i < -n; i++ {
			res /= x
		}
	}
	return res
}

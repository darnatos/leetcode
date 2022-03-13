package solution

func replaceNonCoprimes(nums []int) []int {
	var g int
	res := make([]int, 0)
	for i := range nums {
		res = append(res, nums[i])
		if len(res) < 2 {
			continue
		}

		for len(res) >= 2 {
			g = gcd(res[len(res)-1], res[len(res)-2])
			if g <= 1 {
				break
			}
			res[len(res)-2] = res[len(res)-1] * res[len(res)-2] / g
			res = res[:len(res)-1]
		}
	}
	return res
}

package solution

func Trap(height []int) int {
	lm, rm := height[0], height[len(height)-1]
	l, r := 1, len(height)-2
	res := 0

	for l <= r {
		if lm < rm {
			if height[l] >= lm {
				lm = height[l]
			} else {
				res += lm - height[l]
			}
			l++
		} else {
			if height[r] >= rm {
				rm = height[r]
			} else {
				res += rm - height[r]
			}
			r--
		}
	}
	return res
}

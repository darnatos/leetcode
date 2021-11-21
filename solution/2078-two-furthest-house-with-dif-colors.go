package solution

func MaxDistance(colors []int) int {
	res := 0
	for i := range colors {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] != colors[j] {
				res = max(res, j-i)
			}
		}
	}
	return res
}

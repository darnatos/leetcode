package solution

func GetCollisionTimes(cars [][]int) []float64 {
	res := make([]float64, len(cars))
	stack := make([]int, 0, len(cars))
	for i := len(cars) - 1; i >= 0; i-- {
		res[i] = -1.0
		p1, s1 := float64(cars[i][0]), float64(cars[i][1])
		for len(stack) != 0 {
			j := stack[len(stack)-1]
			p2, s2 := float64(cars[j][0]), float64(cars[j][1])
			if s1 <= s2 || res[j] > 0 && (p2-p1)/(s1-s2) > res[j] {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		if len(stack) != 0 {
			j := stack[len(stack)-1]
			res[i] = (float64(cars[j][0]) - p1) / (s1 - float64(cars[j][1]))
		}
		stack = append(stack, i)
	}
	return res
}

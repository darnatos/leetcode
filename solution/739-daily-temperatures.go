package solution

func DailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	j := -1
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j >= 0 && temperatures[i] > temperatures[stack[j]] {
			res[stack[j]] = i - stack[j]
			j--
		}
		j++
		stack[j] = i
	}
	return res
}

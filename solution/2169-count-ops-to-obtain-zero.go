package solution

func countOperations(num1 int, num2 int) int {
	res := 0
	for num1 > 0 && num2 > 0 {
		if num2 >= num1 {
			num2 -= num1
		} else {
			num1 -= num2
		}
		res++
	}
	return res
}

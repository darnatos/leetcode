package solution

func countEven(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		tmp := 0
		j := i
		for j > 0 {
			tmp += j % 10
			j /= 10
		}
		if tmp%2 == 0 {
			res++
		}
	}
	return res
}

package solution

func Maximum69Number(num int) int {
	arr := make([]int, 0, 5)

	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}

	p6 := 10

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 6 {
			p6 = i
			break
		}
	}

	if p6 != 10 {
		arr[p6] = 9
	}

	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res *= 10
		res += arr[i]
	}
	return res
}

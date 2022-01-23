package solution

func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)

	for i := genNextNum(low); i <= high; i = genNextNum(i + 1) {
		res = append(res, i)
	}

	return res
}

func genNextNum(n int) int {
	stack := make([]int, 0, 9)
	for a := n; a > 0; a /= 10 {
		stack = append(stack, a%10)
	}

	flag := false
	for i := len(stack) - 2; i >= 0; i-- {
		if stack[i] > stack[i+1]+1 {
			flag = true
			break
		} else if stack[i] <= stack[i+1] {
			break
		}
	}

	lm := stack[len(stack)-1]
	if flag {
		lm += 1
	}

	res := 0
	if lm+len(stack) <= 10 {
		for i := lm; i < lm+len(stack); i++ {
			res = 10*res + i
		}
	} else {
		for i := 1; i <= len(stack)+1; i++ {
			res = res*10 + i
		}
	}

	return res
}

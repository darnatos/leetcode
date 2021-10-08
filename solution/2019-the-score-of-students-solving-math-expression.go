package solution

func ScoreOfStudents(s string, answers []int) int {
	n := len(s)/2 + 1
	nums := make([]int, n)
	oprs := make([]byte, n-1)
	for i := 0; i < len(s); i++ {
		if i&1 == 0 {
			nums[i>>1] = int(s[i] - '0')
		} else {
			oprs[i>>1] = s[i]
		}
	}

	dp0 := make([]map[int]bool, n*n)
	dp := make([][]map[int]bool, n)
	for i := range dp {
		dp[i] = dp0[i*n : (i+1)*n]
		dp[i][i] = map[int]bool{nums[i]: true}
	}

	for le := 1; le <= n; le++ {

		for i := 0; i < n-le; i++ {
			j := i + le
			dp[i][j] = make(map[int]bool)
			for k := i; k < j; k++ {
				if oprs[k] == '+' {
					for a := range dp[i][k] {
						for b := range dp[k+1][j] {
							if a+b <= 1000 {
								dp[i][j][a+b] = true
							}
						}
					}
				} else {
					for a := range dp[i][k] {
						for b := range dp[k+1][j] {
							c := a * b
							if c <= 1000 {
								dp[i][j][c] = true
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	ans := calc(s)
	scores := dp[0][n-1]
	for _, v := range answers {
		if v == ans {
			sum += 5
		} else if scores[v] {
			sum += 2
		}
	}
	return sum
}

func calc(s string) int {
	stack := make([]int, len(s)/2)
	num, opr := 0, byte('+')

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = int(ch - '0')
		}
		if i >= len(s)-1 || ch == '+' || ch == '*' {
			if opr == '+' {
				stack = append(stack, num)
			} else if opr == '*' {
				stack[len(stack)-1] = stack[len(stack)-1] * num
			}
			num = 0
			opr = ch
		}
	}

	sum := 0
	for i := range stack {
		sum += stack[i]
	}
	return sum
}

package solution

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}
	mat := [10][10]int{
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
	}

	res := [10]int{}
	for i := range res {
		res[i] = 1
	}
	n--
	for n > 0 {
		if n&1 == 1 {
			tmp := [10]int{}
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					tmp[i] += mat[i][j] * res[j]
				}
				tmp[i] %= 1000000007
			}
			res = tmp
		}

		mm := [10][10]int{}
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					mm[i][j] += mat[i][k] * mat[k][j]
				}
				mm[i][j] %= 1000000007
			}
		}
		mat = mm
		n /= 2
	}

	r := 0
	for i := range res {
		r += res[i]
	}
	return r % 1000000007
}

package solution

func numOfWays(n int) int {
	v := [2]int{6, 6}
	mat := [2][2]int{
		{3, 2},
		{2, 2},
	}
	n--
	for n > 0 {
		if n&1 == 1 {
			v[0], v[1] = (mat[0][0]*v[0]+mat[0][1]*v[1])%mod, (mat[1][0]*v[0]+mat[1][1]*v[1])%mod
		}
		tmp := [2][2]int{}
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					tmp[i][j] += mat[i][k] * mat[k][j]
				}
				tmp[i][j] %= mod
			}
		}
		n >>= 1
		mat = tmp
	}
	return (v[0] + v[1]) % mod
}

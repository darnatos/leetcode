package solution

var minPrimes []int
var comb [10013][14]int

func waysToFillArray(queries [][]int) []int {
	if len(minPrimes) == 0 {
		minPrimes = sievePrimeFactor(10000)
	}
	if comb[0][0] == 0 {
		comb[0][0] = 1
		for i := 1; i < 10013; i++ {
			for j := 0; j < 14; j++ {
				if j == 0 {
					comb[i][j] = 1
				} else {
					comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % mod
				}
			}
		}
	}

	res := make([]int, len(queries))

	for i, q := range queries {
		res[i] = 1
		n, k := q[0], q[1]
		ps := make(map[int]int)
		for k > 1 {
			ps[minPrimes[k]]++
			k /= minPrimes[k]
		}

		for _, f := range ps {
			res[i] = (res[i] * comb[n-1+f][f]) % mod
		}
		// fmt.Println(n,k,ps,res[i])
	}
	return res
}

func sievePrimeFactor(n int) []int {
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		spf[i] = i
	}
	for i := 4; i <= n; i += 2 {
		spf[i] = 2
	}
	for i := 3; i <= n; i++ {
		if spf[i] == i {
			for j := i * i; j <= n; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
	return spf
}

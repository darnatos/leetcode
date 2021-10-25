package solution

func LenLongestFibSubseq1(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	res := 0
	for i := 2; i < n; i++ {
		l, r := 0, i-1
		for l < r {
			sum := arr[l] + arr[r]
			if sum < arr[i] {
				l++
			} else if sum > arr[i] {
				r--
			} else {
				dp[r][i] = dp[l][r] + 1
				res = max(res, dp[r][i])
				l++
				r--
			}
		}
	}
	if res == 0 {
		return 0
	}
	return res + 2
}

func LenLongestFibSubseq2(arr []int) int {
	m := make(map[int]struct{}, len(arr))
	for _, n := range arr {
		m[n] = struct{}{}
	}
	n := len(arr)
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a, b, l := arr[i], arr[j], 2
			for {
				if _, ok := m[a+b]; !ok {
					break
				}
				a, b = b, a+b
				l++
			}
			if res < l {
				res = l
			}
		}
	}
	if res <= 2 {
		return 0
	}
	return res
}

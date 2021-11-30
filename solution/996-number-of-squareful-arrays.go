package solution

import "math"

func NumSquarefulPerms(nums []int) int {
	n := len(nums)
	g := make([][]int, n)

	for i := range nums {
		for j := i + 1; j < n; j++ {
			if isSquare(nums[i] + nums[j]) {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	// Held-Karp Hamiltonian paths DP.
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 1<<n)
		dp[i][1<<i] = 1
	}

	for s := 3; s < (1 << n); s++ {
		for i := 0; i < n; i++ {
			if s&(1<<i) > 0 {
				for _, j := range g[i] {
					if s&(1<<j) > 0 {
						dp[i][s] += dp[j][s^(1<<i)]
					}
				}
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += dp[i][(1<<n)-1]
	}

	for i := 0; i < n; i++ {
		k := 1
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				k++
			}
		}
		res /= k
	}
	return res
}

func isSquare(a int) bool {
	r := math.Floor(math.Sqrt(float64(a)))
	return float64(a) == r*r
}

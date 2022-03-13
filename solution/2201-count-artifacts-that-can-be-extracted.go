package solution

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	dug := make([][]bool, n)
	for i := range dug {
		dug[i] = make([]bool, n)
	}
	for i := range dig {
		dug[dig[i][0]][dig[i][1]] = true
	}

	check := func(art []int) bool {
		for i := art[0]; i <= art[2]; i++ {
			for j := art[1]; j <= art[3]; j++ {
				if !dug[i][j] {
					return false
				}
			}
		}
		return true
	}

	res := 0
	for i := range artifacts {
		if check(artifacts[i]) {
			res++
		}
	}
	return res
}

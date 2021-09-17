package solution
func HighestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	q := make([][2]int, 0, 8)
	for i := range isWater {
		for j := range isWater[0] {
			if isWater[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	tmp := make([]int, m*n)
	for i := range tmp {
		tmp[i] = -1
	}
	res := make([][]int, m)
	for i := range res {
		res[i] = tmp[i*n:i*n+n]
	}

	k := 0
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[i]
			x, y := cur[0], cur[1]
			if res[x][y] == -1 {
				res[x][y] = k
				if x > 0 {
					q = append(q, [2]int{x-1, y})
				}
				if x < m-1 {
					q = append(q, [2]int{x+1, y})
				}
				if y > 0 {
					q = append(q, [2]int{x, y-1})
				}
				if y < n-1 {
					q = append(q, [2]int{x, y+1})
				}
			}
		}
		q = q[l:]
		k++
	}

	return res
}

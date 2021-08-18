package solution

func NearestExit(maze [][]byte, entrance []int) int {
	maze[entrance[0]][entrance[1]] = '+'
	q := make([][]int, 0, 16)
	q = append(q, []int{entrance[0], entrance[1]})

	res := 0
	for len(q) != 0 {
		for size := len(q); size > 0; size-- {
			tmp := q[0]
			q = q[1:]

			d := [5]int{-1, 0, 1, 0, -1}
			for i := 0; i < 4; i++ {
				r := tmp[0] + d[i]
				c := tmp[1] + d[i+1]
				isOutside := outside(maze, r, c)
				if isOutside && !(tmp[0] == entrance[0] && tmp[1] == entrance[1]) {
					return res
				}
				if !isOutside && maze[r][c] == '.' {
					q = append(q, []int{r, c})
					maze[r][c] = '+'
				}
			}
		}
		res++
	}
	return -1
}

func outside(maze [][]byte, r, c int) bool {
	return r < 0 || r >= len(maze) || c < 0 || c >= len(maze[0])
}

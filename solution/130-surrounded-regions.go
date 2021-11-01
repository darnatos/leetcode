package solution

func Solve(board [][]byte) {
	m, n := len(board), len(board[0])
	dir := []int{0, 1, 0, -1, 0}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if visited[i][j] || board[i][j] == 'X' {
				continue
			}
			sDfs(board, i, j, m, n, dir, visited)
		}
	}
}

func sDfs(board [][]byte, i, j, m, n int, dir []int, visited [][]bool) {
	visited[i][j] = true

	shouldFlip := true
	queue := make([][]int, 0, 4)
	queue = append(queue, []int{i, j})
	q := 0
	for q < len(queue) {
		cur := queue[q]
		q++
		x0, y0 := cur[0], cur[1]
		if !shouldFlip || x0 == 0 || y0 == 0 || x0 == m-1 || y0 == n-1 {
			shouldFlip = false
		}
		for k := 0; k < 4; k++ {
			x, y := x0+dir[k], y0+dir[k+1]
			if x < 0 || y < 0 || x > m-1 || y > n-1 || visited[x][y] || board[x][y] == 'X' {
				continue
			}
			visited[x][y] = true
			queue = append(queue, []int{x, y})
		}
	}
	if shouldFlip {
		for _, pos := range queue {
			board[pos[0]][pos[1]] = 'X'
		}
	}
}

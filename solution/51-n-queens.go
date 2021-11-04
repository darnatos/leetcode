package solution

func SolveNQueens(n int) [][]string {
	tmp := make([]bool, n*n)
	board := make([][]bool, n)
	for i := range board {
		board[i] = tmp[i*n : i*n+n]
	}
	res := make([][]string, 0)

	nqDfs(board, 0, n, &res)
	return res
}

func nqDfs(board [][]bool, y, n int, res *[][]string) {
	if y == n {
		*res = append(*res, construct(board, n))
		return
	}
	for i := 0; i < n; i++ {
		if validate(board, i, y, n) {
			board[i][y] = true
			nqDfs(board, y+1, n, res)
			board[i][y] = false
		}
	}
}

func validate(board [][]bool, x, y, n int) bool {
	for j := 0; j < n; j++ {
		if board[x][j] {
			return false
		}
	}
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	for i, j := x+1, y-1; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}
	return true
}

func construct(board [][]bool, n int) []string {
	res := make([]string, n)
	for i := range board {
		str := make([]byte, n)
		for j := range board[i] {
			if board[i][j] {
				str[j] = 'Q'
			} else {
				str[j] = '.'
			}
		}
		res[i] = string(str)
	}
	return res
}

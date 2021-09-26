package solution

func PlaceWordInCrossword(board [][]byte, word string) bool {
	m, n, w := len(board), len(board[0]), len(word)
	if w <= n {
		for i := range board {
			for j := 0; j <= n-w; j++ {
				if j > 0 && board[i][j-1] != '#' || j < n-w && board[i][j+w] != '#' {
					continue
				}
				k := 0
				for ; k < w; k++ {
					if board[i][j+k] == ' ' || board[i][j+k] == word[k] {
						continue
					}
					break
				}
				if k == w {
					return true
				}
				k = 0
				for ; k < w; k++ {
					if board[i][j+k] == ' ' || board[i][j+k] == word[w-1-k] {
						continue
					}
					break
				}
				if k == w {
					return true
				}
			}
		}
	}
	if w <= m {
		for j := range board[0] {
			for i := 0; i <= m-w; i++ {
				if i > 0 && board[i-1][j] != '#' || i < m-w && board[i+w][j] != '#' {
					continue
				}
				k := 0
				for ; k < w; k++ {
					if board[i+k][j] == ' ' || board[i+k][j] == word[k] {
						continue
					}
					break
				}
				if k == w {
					return true
				}
				k = 0
				for ; k < w; k++ {
					if board[i+k][j] == ' ' || board[i+k][j] == word[w-1-k] {
						continue
					}
					break
				}
				if k == w {
					return true
				}
			}
		}
	}
	return false
}
